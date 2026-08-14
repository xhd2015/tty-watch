# Scenario

**Feature**: idle detached session reports zero watch and attached clients

```
# detached sleep, no WS clients
harness -> detached sleep -> tty-watch list -> WATCH=0 ATTACHED=0
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-table-idle"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```