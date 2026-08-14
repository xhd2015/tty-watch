# Scenario

**Feature**: attach connects to a detached session and streams live PTY output

```
# detached echo loop; attach observer receives live bytes
harness -> run sleep (detached) -> tty-watch attach <id> -> ATTACH_LIVE_MARKER on stdout
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-detached-session"
	req.AttachProbe = "2s"
	return nil
}
```