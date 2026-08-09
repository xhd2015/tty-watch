package ttywatch

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/xhd2015/dot-pkgs/go-pkgs/shell/ptywrap"
)

// ServeOptions configures an embedded PTY server session.
type ServeOptions struct {
	SessionID      string
	Command        []string
	Home           string
	RegistrySubdir string
	Cwd            string
	ExtraPaths     []string
	KeepAlive      bool
	// CommandEnv is KEY=VALUE assignments merged into the PTY agent child environ.
	CommandEnv []string
	// CommandUnset is env keys removed from the PTY agent child environ before CommandEnv.
	CommandUnset []string
	// OnListening is invoked in a goroutine after the PTY session is created and
	// the registry entry is written. ctx is cancelled when the serve session ends.
	// listenAddr is the bound HTTP listen address; home and registrySubdir are the
	// resolved registry location for this serve process.
	OnListening func(ctx context.Context, listenAddr, home, registrySubdir string)
}

// serveKeepAlive reports keep-alive from explicit opts only (policy B: no ambient
// TTY_WATCH_KEEP_ALIVE).
func serveKeepAlive(opts ServeOptions) bool {
	return opts.KeepAlive
}

// ServeSession runs the embedded ptywrap HTTP server until the command exits.
func ServeSession(ctx context.Context, opts ServeOptions) error {
	if opts.SessionID == "" {
		return fmt.Errorf("serve: missing session id")
	}
	if len(opts.Command) == 0 {
		return fmt.Errorf("serve: missing command")
	}

	home := opts.Home
	if home == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		home = DefaultTTYWatchHome(userHome)
	}
	cwd := opts.Cwd
	if cwd == "" {
		var err error
		cwd, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	mgr := ptywrap.NewManager()
	if len(opts.ExtraPaths) > 0 {
		mgr.Spawn.ExtraPaths = append([]string(nil), opts.ExtraPaths...)
	}
	if len(opts.CommandEnv) > 0 {
		mgr.Spawn.Env = append([]string(nil), opts.CommandEnv...)
	}
	if len(opts.CommandUnset) > 0 {
		mgr.Spawn.Unset = append([]string(nil), opts.CommandUnset...)
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	listenAddr := ln.Addr().String()

	mux := http.NewServeMux()
	registerPrepareInjectAPI(mux, mgr)
	ptywrap.RegisterAPIWithManager(mux, mgr)
	srv := &http.Server{Handler: mux}
	serverDone := make(chan error, 1)
	go func() {
		serverDone <- srv.Serve(ln)
	}()

	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
		<-serverDone
	}

	if _, err := mgr.CreateCommandWithID(opts.SessionID, "tty-watch", cwd, opts.Command); err != nil {
		shutdown()
		return err
	}

	// Resolve PTY agent child PID (direct child of this __serve__ process).
	// Used by lifecycle probes so keep-alive serve PID is not mistaken for "agent live".
	commandPID := 0
	if child, err := firstChildPID(os.Getpid()); err == nil && child > 0 {
		commandPID = child
	}

	entry := RegistryEntry{
		SessionID:  opts.SessionID,
		ListenAddr: listenAddr,
		PID:        os.Getpid(),
		CreatedAt:  time.Now().UTC().Format(time.RFC3339),
		Command:    opts.Command,
		Cwd:        cwd,
		CommandPID: commandPID,
	}
	cfg := serveRegistryConfig(home, opts.RegistrySubdir)
	if err := WriteRegistry(cfg, entry); err != nil {
		mgr.Remove(opts.SessionID)
		shutdown()
		return err
	}

	// Session-scoped cancel so OnListening consumers (e.g. send-queue drainer)
	// stop when this serve exits, even if the parent ctx is never cancelled.
	serveCtx, serveCancel := context.WithCancel(ctx)
	defer serveCancel()

	if opts.OnListening != nil {
		go opts.OnListening(serveCtx, listenAddr, home, cfg.Subdir)
	}

	waitDone := make(chan struct{})
	go func() {
		_ = mgr.Wait(opts.SessionID)
		close(waitDone)
	}()

	select {
	case <-ctx.Done():
		serveCancel()
		mgr.Remove(opts.SessionID)
		shutdown()
		RemoveRegistryIfMatch(cfg, opts.SessionID, listenAddr, entry.PID)
		return ctx.Err()
	case <-waitDone:
	}

	// Record durable agent-exit so probes need not re-scan process tables.
	entry.CommandExited = true
	entry.CommandExitedAt = time.Now().UTC().Format(time.RFC3339)
	_ = WriteRegistry(cfg, entry)

	if serveKeepAlive(opts) {
		// Keep the ptywrap server and registry reachable after the PTY child exits
		// so web/CLI attach can replay scrollback on finished keep-tty sessions.
		<-ctx.Done()
		serveCancel()
		shutdown()
		mgr.Remove(opts.SessionID)
		RemoveRegistryIfMatch(cfg, opts.SessionID, listenAddr, entry.PID)
		return ctx.Err()
	}

	serveCancel()
	const writerAttachGrace = 2 * time.Second
	time.Sleep(writerAttachGrace)
	shutdown()
	mgr.Remove(opts.SessionID)
	RemoveRegistryIfMatch(cfg, opts.SessionID, listenAddr, entry.PID)
	return nil
}