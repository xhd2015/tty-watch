package cli

import (
	"fmt"

	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func runAttach(cfg Config) error {
	if cfg.Attach == nil || cfg.Attach.Session == "" {
		return fmt.Errorf("attach: requires <session-id>")
	}
	sessionID := cfg.Attach.Session

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

	_, err = ttywatch.AttachWriter(entry.ListenAddr, sessionID, "attach")
	return err
}
