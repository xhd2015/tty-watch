# Scenario

**Feature**: unknown serve flag is an error

```
Args [sid --not-a-real-flag -- echo hi] -> ParseServeArgv -> error
```

## Steps

1. Insert an unrecognized long flag before `--` and a valid command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"bad-sess", "--not-a-real-flag", "--", "echo", "hi"}
	return nil
}
```
