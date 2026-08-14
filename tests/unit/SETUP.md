# Scenario

**Feature**: tty-watch unit-level conversion helpers + pure EncodeSGRClick

```
scrollback bytes -> screen snapshot frame -> plain text
EncodeSGRClick(row,col,btn,release) -> CSI SGR mouse wire bytes (no session)
source-order checks on pkgs/ttywatch/*.go (observer/attach)
```

## Preconditions

- Root Setup builds `./cmd/tty-watch` (unused by pure encode; used for module-root resolution in source checks).
- `encode-sgr-click/` uses **Mode=encode** (in-process `ttywatch.EncodeSGRClick`).
- Other unit leaves set `req.Phase` (e.g. `unit-normalize-tty-output`, source checks).

## Steps

1. encode subtree sets Mode=encode + Encode* fields.
2. Other leaves set Phase for harness unit phases / source checks.
3. Assert checks wire bytes, column-zero layout, or source-order flags.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Unit leaves are pure helpers / source checks: no live session required.
	// encode-sgr-click sets Mode=encode; other leaves set Phase in their own Setup.
	if req.Mode == "" && req.Phase == "" {
		// Default to phase path; leaf must override Phase (or Mode for encode).
		req.Phase = "unit-pending"
	}
	return nil
}
```
