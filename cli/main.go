package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

// Main is the programmatic entry point for the tty-watch CLI.
// args is argv without the program name (same as os.Args[1:]).
// Returns an error on failure; does not os.Exit and does not print "Error:" prefix.
func Main(args []string, stdin io.Reader, stdout, stderr io.Writer) error {
	if stdin == nil {
		stdin = os.Stdin
	}
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}
	cfg, err := ParseArgs(args)
	if err != nil {
		return err
	}
	cfg.Stdin = stdin
	cfg.Stdout = stdout
	cfg.Stderr = stderr
	return Run(cfg)
}

// ParseArgs maps argv (no program name) to a fully populated Config.
// Flag/mode validation only; does not touch registry or PTY.
func ParseArgs(args []string) (Config, error) {
	var cfg Config
	if len(args) == 0 {
		cfg.Command = "help"
		return cfg, nil
	}

	if ttywatch.IsServeSubcommand(args[0]) {
		parsed, err := ttywatch.ParseServeArgv(args[1:])
		if err != nil {
			return cfg, err
		}
		cfg.Command = args[0]
		cfg.Serve = &ServeOptions{
			SessionID:      parsed.SessionID,
			Command:        append([]string(nil), parsed.Command...),
			Home:           parsed.Home,
			RegistrySubdir: parsed.RegistrySubdir,
			KeepAlive:      parsed.KeepAlive,
			ExtraPaths:     append([]string(nil), parsed.ExtraPaths...),
			CommandEnv:     append([]string(nil), parsed.CommandEnv...),
			CommandUnset:   append([]string(nil), parsed.CommandUnset...),
		}
		return cfg, nil
	}

	switch args[0] {
	case "run":
		opts, err := parseRunArgs(args[1:])
		if err != nil {
			return cfg, err
		}
		cfg.Command = "run"
		cfg.Run = &RunOptions{
			Headless:  opts.headless,
			Detach:    opts.detach,
			SessionID: opts.customSessionID,
			Command:   opts.commandArgs,
		}
		return cfg, nil

	case "list":
		if len(args) > 1 {
			return cfg, fmt.Errorf("list: unexpected arguments %v", args[1:])
		}
		cfg.Command = "list"
		cfg.List = &ListOptions{}
		return cfg, nil

	case "watch":
		if len(args) != 2 {
			return cfg, fmt.Errorf("watch: requires <session-id>")
		}
		cfg.Command = "watch"
		cfg.Watch = &WatchOptions{Session: args[1]}
		return cfg, nil

	case "attach":
		if len(args) != 2 {
			return cfg, fmt.Errorf("attach: requires <session-id>")
		}
		cfg.Command = "attach"
		cfg.Attach = &AttachOptions{Session: args[1]}
		return cfg, nil

	case "snapshot":
		if len(args) != 2 {
			return cfg, fmt.Errorf("snapshot: requires <session-id>")
		}
		cfg.Command = "snapshot"
		cfg.Snapshot = &SnapshotOptions{Session: args[1]}
		return cfg, nil

	case "kill":
		if len(args) != 2 {
			return cfg, fmt.Errorf("kill: requires <session-id>")
		}
		cfg.Command = "kill"
		cfg.Kill = &KillOptions{Session: args[1]}
		return cfg, nil

	case "send":
		opts, err := parseSendArgs(args[1:])
		if err != nil {
			return cfg, err
		}
		cfg.Command = "send"
		cfg.Send = &opts
		return cfg, nil

	case "-h", "--help", "help":
		cfg.Command = "help"
		return cfg, nil

	default:
		return cfg, fmt.Errorf("unknown subcommand %q", args[0])
	}
}

// Run executes a ready Config. Streams come from cfg (no package-level globals).
// Does not parse flags and does not os.Exit.
func Run(cfg Config) error {
	if cfg.Stdout == nil {
		cfg.Stdout = os.Stdout
	}
	if cfg.Stderr == nil {
		cfg.Stderr = os.Stderr
	}
	if cfg.Stdin == nil {
		cfg.Stdin = os.Stdin
	}

	// Internal serve reexec token (slug form).
	if ttywatch.IsServeSubcommand(cfg.Command) {
		return runServe(cfg)
	}

	switch cfg.Command {
	case "", "help":
		printHelp(cfg.Stdout)
		return nil
	case "run":
		return runRun(cfg)
	case "list":
		return runList(cfg)
	case "watch":
		return runWatch(cfg)
	case "attach":
		return runAttach(cfg)
	case "snapshot":
		return runSnapshot(cfg)
	case "kill":
		return runKill(cfg)
	case "send":
		return runSend(cfg)
	case "serve":
		return runServe(cfg)
	default:
		return fmt.Errorf("unknown subcommand %q", cfg.Command)
	}
}

func resolveHome(cfg Config) (string, error) {
	if cfg.Home != "" {
		return cfg.Home, nil
	}
	return TTYWatchHome()
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, `tty-watch — embedded PTY session manager

Usage:
  tty-watch run [--headless] [--detach] [--session-id <id>] <command> [args...]
                                     Start session; default attaches (writer)
                                     --headless prints session-id and waits
                                     --detach prints session-id and exits
  tty-watch list                     List sessions
  tty-watch watch <session-id>       Observe session (readonly)
  tty-watch attach <session-id>      Join session (write+resize)
  tty-watch snapshot <session-id>    Print sanitized scrollback
  tty-watch kill <session-id>        End session and remove registry
  tty-watch send <session-id> <msg>...
                                     Inject follow-up text into live session
  tty-watch send <session-id> --click --row r --col c
                 [--mouse btn] [--no-release] [--json]
                                     Inject SGR mouse click (0-based row/col)
  tty-watch send <session-id> --query-cursor [--json]
                                     Print host VT cursor (0-based; no child inject)

Options:
  -h, --help                         Show help
`)
}
