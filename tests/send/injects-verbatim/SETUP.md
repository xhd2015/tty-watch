# Scenario

**Feature**: send injects multi-word message bytes exactly as provided

```
# detached cat capture; send follow-up text
harness -> detached cat capture -> tty-watch send session-N "follow up" -> capture.bin == follow up
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "send-injects-verbatim"
	req.SendMessage = "follow up"
	return nil
}
```