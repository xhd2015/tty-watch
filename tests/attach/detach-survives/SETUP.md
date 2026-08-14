# Scenario

**Feature**: attach Ctrl-] (`\x1d`) detaches client only; session survives

```
# attach then detach; registry + list still show session
harness -> detached sleep -> tty-watch attach -> \x1d -> exit 0; list includes session
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-detach-survives"
	req.RunCommand = []string{"sleep", "300"}
	return nil
}
```