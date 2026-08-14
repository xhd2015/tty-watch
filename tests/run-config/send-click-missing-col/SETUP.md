# Scenario

**Feature**: `Run` rejects Send click options with missing/invalid col (col &lt; 0)

```
# programmatic path — no argv; validation inside Run/SendOptions
cli.Run(Config{
  Command: "send",
  Send: &SendOptions{Session: "sess-flags", Mode: SendModeClick, Row: 1, Col: -1},
}) -> error about col
# does not require registry (validation before inject)
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"

	"github.com/xhd2015/tty-watch/cli"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "run-config"
	req.Config = cli.Config{
		Command: "send",
		Send: &cli.SendOptions{
			Session: "sess-flags",
			Mode:    cli.SendModeClick,
			Row:     1,
			Col:     -1, // missing/invalid: Run requires col ≥ 0 for click
		},
	}
	return nil
}
```
