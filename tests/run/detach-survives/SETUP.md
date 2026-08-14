# Scenario

**Feature**: Ctrl-] (`\x1d`) detaches client; session survives in registry and list

```
# detach client only
harness PTY -> tty-watch sleep -> \x1d -> list still shows session
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-detach"
	req.Detach = true
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```