# Scenario

**Feature**: headless accepts `--session-id` before command argv

```
# custom id printed on stdout and written to registry
harness pipe -> tty-watch run --headless --session-id my-job sleep 120
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-headless-with-custom-session-id"
	req.CustomSessionID = "my-job"
	req.RunCommand = []string{"sleep", "120"}
	return nil
}
```