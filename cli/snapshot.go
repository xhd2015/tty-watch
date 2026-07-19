package cli

import "fmt"

func runSnapshot(cfg Config) error {
	if cfg.Snapshot == nil || cfg.Snapshot.Session == "" {
		return fmt.Errorf("snapshot: requires <session-id>")
	}
	sessionID := cfg.Snapshot.Session

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
		return fmt.Errorf("tty-watch session %s not found", sessionID)
	}

	frame, scrollback, cols, rows, err := readSnapshot(entry.ListenAddr, sessionID)
	if err != nil {
		return err
	}
	text := renderSnapshotOutput(frame, scrollback, cols, rows)
	if text != "" {
		fmt.Fprintln(cfg.Stdout, text)
	}
	return nil
}
