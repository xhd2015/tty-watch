package cli

import "io"

// Config is the fully config-driven input for Run. No argv parsing inside Run.
type Config struct {
	Command string // "run"|"list"|"watch"|"attach"|"snapshot"|"kill"|"send"|"help"|""|serve token
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Home    string // optional TTY_WATCH_HOME override

	Run      *RunOptions
	List     *ListOptions
	Watch    *WatchOptions
	Attach   *AttachOptions
	Snapshot *SnapshotOptions
	Kill     *KillOptions
	Send     *SendOptions
	Serve    *ServeOptions
}

// RunOptions holds flags for the run subcommand.
type RunOptions struct {
	Headless  bool
	Detach    bool
	SessionID string
	Command   []string
}

// ListOptions is reserved for list (currently no flags).
type ListOptions struct{}

// WatchOptions holds the session id for watch.
type WatchOptions struct {
	Session string
}

// AttachOptions holds the session id for attach.
type AttachOptions struct {
	Session string
}

// SnapshotOptions holds the session id for snapshot.
type SnapshotOptions struct {
	Session string
}

// KillOptions holds the session id for kill.
type KillOptions struct {
	Session string
}

// ServeOptions holds internal reexec serve parameters.
type ServeOptions struct {
	SessionID string
	Command   []string
}

// SendMode is the exclusive mode for tty-watch send.
type SendMode int

const (
	SendModeText SendMode = iota
	SendModeClick
	SendModeQueryCursor
)

// SendOptions holds validated send parameters (text | click | query-cursor).
type SendOptions struct {
	Session   string
	Mode      SendMode
	Message   string
	Row, Col  int // 0-based click coordinates
	Mouse     int
	NoRelease bool // default release ON when false
	JSON      bool
}
