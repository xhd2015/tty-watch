package cli

import (
	"fmt"

	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func runWatch(cfg Config) error {
	if cfg.Watch == nil || cfg.Watch.Session == "" {
		return fmt.Errorf("watch: requires <session-id>")
	}
	sessionID := cfg.Watch.Session

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

	return ttywatch.StreamObserver(entry.ListenAddr, sessionID, cfg.Stdout)
}
