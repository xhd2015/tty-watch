# Scenario

**Bug**: `echo yes` on interactive TTY emitted LF-only snapshot text, smearing `[Terminal exited]` right

```
tty-watch run echo yes -> raw PTY bytes:
  yes\r\n[Terminal exited]\r\n
(without CR, host prompt renders far right after exit)
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-bash-c-exit-marker-column-zero"
	req.RunCommand = []string{"echo", "yes"}
	return nil
}
```