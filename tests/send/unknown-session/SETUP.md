# Scenario

**Feature**: send on unknown session id fails when registry is empty

```
# missing registry entry
tty-watch send session-999 "hi" -> error: session not found
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "send-missing"
	req.SendID = "session-99999"
	req.SendMessage = "hi"
	return nil
}
```