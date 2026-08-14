# Scenario

**Feature**: headless blocks until inner command exits naturally

```
# short command exits; headless parent follows with exit 0
harness pipe -> tty-watch run --headless true -> wait -> registry pruned
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-headless-waits-until-child-exits"
	req.RunCommand = []string{"true"}
	return nil
}
```