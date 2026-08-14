# Scenario

**Feature**: attach on unknown session id fails

```
# missing registry entry
tty-watch attach session-99999 -> exit 1; not found error
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-unknown"
	req.AttachID = "session-99999"
	return nil
}
```