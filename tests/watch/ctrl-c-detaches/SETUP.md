# Scenario

**Feature**: Ctrl-C detaches watch without killing the remote session

```
# detached sleep session; watch on PTY; harness sends \x03; watch exits; session survives
harness PTY -> watch -> Ctrl-C -> watch detaches; registry entry remains
```

Bug: `streamObserver` does not read stdin or handle Ctrl-C, so watch stays attached
when the user presses Ctrl-C. Expected: Ctrl-C cleanly detaches the observer;
the remote session keeps running.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "watch-ctrl-c-detaches"
	return nil
}
```