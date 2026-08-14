# Scenario

**Feature**: `tty-watch list` prints an aligned table header when sessions exist

```
# detached sleep session, no clients
harness -> detached sleep -> tty-watch list -> SESSION UPTIME WATCH ATTACHED COMMAND header + aligned row
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-table-header"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```