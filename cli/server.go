package cli

import (
	"context"
	"fmt"

	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func runServe(cfg Config) error {
	if cfg.Serve == nil {
		return fmt.Errorf("serve: missing session id or command")
	}
	if cfg.Serve.SessionID == "" || len(cfg.Serve.Command) == 0 {
		return fmt.Errorf("serve: missing session id or command")
	}
	return ttywatch.ServeSession(context.Background(), ttywatch.ServeOptions{
		SessionID: cfg.Serve.SessionID,
		Command:   cfg.Serve.Command,
	})
}
