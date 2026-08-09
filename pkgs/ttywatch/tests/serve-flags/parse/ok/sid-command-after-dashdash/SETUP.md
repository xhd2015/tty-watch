# Scenario

**Feature**: minimal parse — sid and command after `--`, no serve flags

```
Args [sess-1 -- echo hello] -> ParseServeArgv
  -> SessionID=sess-1, Command=[echo hello], empty knobs
```

## Steps

1. Set `req.Args` to session id, `--`, and a two-word command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"sess-1", "--", "echo", "hello"}
	return nil
}
```
