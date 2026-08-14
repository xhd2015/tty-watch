# Scenario

**Feature**: `Main([]string{"-h"})` prints root usage and returns nil error

```
cli.Main(["-h"], …) -> nil error; stdout Usage + subcommands (run, send, …)
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"-h"}
	return nil
}
```
