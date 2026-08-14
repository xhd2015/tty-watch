# Scenario

**Feature**: invalid custom session ids are rejected before starting a session

```
# leading dot violates validation pattern
tty-watch run --session-id .bad sleep 1 -> exit 1
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-custom-invalid-id"
	req.CustomSessionID = ".bad"
	req.RunCommand = []string{"sleep", "1"}
	return nil
}
```