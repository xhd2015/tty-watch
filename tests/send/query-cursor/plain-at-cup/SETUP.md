# Scenario

**Feature**: plain `--query-cursor` prints 0-based cursor after CUP 5;3

```
detached printf '\033[5;3H'; sleep -> tty-watch send <id> --query-cursor
  -> stdout "row=4 col=2\n"
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Plain stdout (not JSON); CUP fixture is default in phaseSendQueryCursor.
	req.JSON = false
	return nil
}
```
