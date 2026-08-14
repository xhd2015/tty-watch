# Scenario

**Feature**: list with empty registry produces no session lines

```
# fresh home, no sessions
tty-watch list -> empty registry -> no session-N lines
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-empty"
	return nil
}
```