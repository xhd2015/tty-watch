# Scenario

**Feature**: list prints session id, joined command argv, and human uptime

```
# detached sleep session
harness -> detached sleep -> tty-watch list -> id + sleep + uptime
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-fields"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```