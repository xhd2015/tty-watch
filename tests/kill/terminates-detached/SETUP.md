# Scenario

**Feature**: kill terminates a detached session and removes registry entry

```
# detached sleep then kill
harness -> detached sleep -> tty-watch kill -> registry gone, not reachable
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "kill-stop"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```