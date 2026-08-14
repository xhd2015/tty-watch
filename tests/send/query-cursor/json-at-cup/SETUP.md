# Scenario

**Feature**: `--query-cursor --json` prints JSON cursor after CUP 5;3

```
tty-watch send <id> --query-cursor --json -> {"row":4,"col":2}
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.JSON = true
	return nil
}
```
