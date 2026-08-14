# Scenario

**Feature**: short-lived `echo` prints output and exits promptly

```
harness PTY -> tty-watch run echo yes -> prints yes -> exit 0 within 8s
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-echo-exits"
	req.RunCommand = []string{"echo", "yes"}
	return nil
}
```