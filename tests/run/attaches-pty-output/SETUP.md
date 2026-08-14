# Scenario

**Feature**: writer attach bridges PTY output from child command

```
# echo marker visible on harness PTY before detach
harness PTY -> tty-watch sh -c 'echo RUN_OK' -> RUN_OK on PTY capture
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-attach-output"
	req.RunCommand = []string{"sh", "-c", "echo RUN_OK; sleep 60"}
	return nil
}
```