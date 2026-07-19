package cli

import (
	"fmt"
	"syscall"
	"time"

	ptyclient "github.com/xhd2015/dot-pkgs/go-pkgs/shell/ptywrap/client"
)

func runKill(cfg Config) error {
	if cfg.Kill == nil || cfg.Kill.Session == "" {
		return fmt.Errorf("kill: requires <session-id>")
	}
	sessionID := cfg.Kill.Session

	home, err := resolveHome(cfg)
	if err != nil {
		return err
	}
	entry, err := ReadRegistry(home, sessionID)
	if err != nil {
		return err
	}

	if !tcpReachable(entry.ListenAddr) {
		RemoveRegistryIfMatch(home, sessionID, entry.ListenAddr, entry.PID)
		return nil
	}

	c := ptyclient.NewClient("http://" + entry.ListenAddr)
	_ = c.Delete(sessionID)

	if entry.PID > 0 && processAlive(entry.PID) {
		_ = syscall.Kill(entry.PID, syscall.SIGTERM)
		deadline := time.Now().Add(2 * time.Second)
		for time.Now().Before(deadline) && processAlive(entry.PID) {
			time.Sleep(100 * time.Millisecond)
		}
		if processAlive(entry.PID) {
			_ = syscall.Kill(entry.PID, syscall.SIGKILL)
		}
	}

	RemoveRegistryIfMatch(home, sessionID, entry.ListenAddr, entry.PID)
	return nil
}
