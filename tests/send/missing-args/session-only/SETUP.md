# Scenario

**Feature**: `send` with session id but no message/mode is an error

```
cli.Main(["send", "sess-flags"], …) -> error; requires message
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags"}
	return nil
}
```
