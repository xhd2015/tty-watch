# Scenario

**Feature**: ATTACHED includes screen writer connection (`tty-watch run` attach)

```
# detached sleep + screen-mode WS writer (same role as run attach)
harness -> detached sleep -> dial screen writer -> tty-watch list -> ATTACHED>=1
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-table-writer"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```