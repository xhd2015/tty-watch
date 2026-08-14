# Scenario

**Feature**: successful click injects SGR bytes into PTY (cat capture)

```
# detached cat > capture.bin; send --click ...
harness -> detached cat capture -> tty-watch send --click --row R --col C
  -> capture.bin == SGR press [+ release]
```

## Preconditions

- Capture child: `cat > inject-capture.bin; sleep 300` (harness `byteCaptureSessionCommand` via `StartDetachedSession`).
- **Requires working `run --detach` / serve.** If serve is SIGKILL'd in the environment, harness times out on registry wait (infra, not product assert).
- Assert uses `resp.InjectedBytes` for wire equality.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "send-click-capture"
	return nil
}
```
