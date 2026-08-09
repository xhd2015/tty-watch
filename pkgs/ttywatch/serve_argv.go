package ttywatch

import (
	"fmt"
	"path/filepath"

	lessflags "github.com/xhd2015/less-flags"
)

// ServeArgv is the parsed form of serve re-exec argv after the serve token:
//
//	<sid> [flags…] [--] <command…>
type ServeArgv struct {
	SessionID      string
	Home           string
	RegistrySubdir string
	KeepAlive      bool
	ExtraPaths     []string
	CommandEnv     []string // KEY=VALUE
	CommandUnset   []string // KEY
	Command        []string
}

// ParseServeArgv parses argv after the serve token into ServeArgv.
// Wire: <sid> [flags…] -- <command…>, or bare command without -- when the
// first remaining token is not a serve flag.
func ParseServeArgv(args []string) (ServeArgv, error) {
	var out ServeArgv
	if len(args) == 0 {
		return out, fmt.Errorf("serve: missing session id")
	}
	out.SessionID = args[0]

	remain, err := lessflags.String("--home", &out.Home).
		String("--registry-subdir", &out.RegistrySubdir).
		Bool("--keep-alive", &out.KeepAlive).
		StringSlice("--extra-path", &out.ExtraPaths).
		StringSlice("--env", &out.CommandEnv).
		StringSlice("--unset-env", &out.CommandUnset).
		StopOnFirstArg().
		Parse(args[1:])
	if err != nil {
		return out, err
	}
	if len(remain) == 0 {
		return out, fmt.Errorf("serve: missing command")
	}
	out.Command = append([]string(nil), remain...)
	return out, nil
}

// BuildServeArgv builds the re-exec remainder after host binary + serve token:
//
//	<sid> [flags…] -- <command…>
//
// Flag emit order: --home, --registry-subdir, --keep-alive, --extra-path…,
// --env…, --unset-env…, then --, then pure command. Empty knobs are omitted.
func BuildServeArgv(sessionID string, opts HeadlessRunOptions) []string {
	out := make([]string, 0, 8+len(opts.ExtraPaths)*2+len(opts.CommandEnv)*2+len(opts.CommandUnset)*2+len(opts.Command))
	out = append(out, sessionID)
	if opts.Home != "" {
		out = append(out, "--home", opts.Home)
	}
	if opts.RegistrySubdir != "" {
		out = append(out, "--registry-subdir", opts.RegistrySubdir)
	}
	if opts.KeepAlive {
		out = append(out, "--keep-alive")
	}
	for _, p := range opts.ExtraPaths {
		out = append(out, "--extra-path", p)
	}
	for _, e := range opts.CommandEnv {
		out = append(out, "--env", e)
	}
	for _, u := range opts.CommandUnset {
		out = append(out, "--unset-env", u)
	}
	out = append(out, "--")
	out = append(out, opts.Command...)
	return out
}

// BuildServeChildEnv returns the re-exec child environ. It must not invent or
// clear TTY_WATCH_HOME, TTY_WATCH_REGISTRY_SUBDIR, TTY_WATCH_KEEP_ALIVE, or
// TTY_WATCH_EXTRA_PATHS — those knobs travel only as serve flags.
func BuildServeChildEnv(base []string, opts HeadlessRunOptions) []string {
	_ = opts
	if base == nil {
		return nil
	}
	return append([]string(nil), base...)
}

// DefaultTTYWatchHome joins userHome with ".tty-watch". Pure: no ambient env.
func DefaultTTYWatchHome(userHome string) string {
	return filepath.Join(userHome, defaultHomeDir)
}

// DefaultRegistrySubdir returns the constant default registry subdirectory.
func DefaultRegistrySubdir() string {
	return defaultSubdir
}

// ServeOptionsFromArgv maps a parsed ServeArgv onto ServeOptions for spawn.
func ServeOptionsFromArgv(a ServeArgv) ServeOptions {
	return ServeOptions{
		SessionID:      a.SessionID,
		Home:           a.Home,
		RegistrySubdir: a.RegistrySubdir,
		KeepAlive:      a.KeepAlive,
		ExtraPaths:     append([]string(nil), a.ExtraPaths...),
		CommandEnv:     append([]string(nil), a.CommandEnv...),
		CommandUnset:   append([]string(nil), a.CommandUnset...),
		Command:        append([]string(nil), a.Command...),
	}
}
