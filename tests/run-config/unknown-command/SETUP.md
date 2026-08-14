# Scenario

**Feature**: `Run(Config{Command:"nope"})` returns a non-nil error

```
cli.Run(Config{Command: "nope", …}) -> error
# message: unknown command/subcommand
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"

	"github.com/xhd2015/tty-watch/cli"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "run-config"
	req.Config = cli.Config{
		Command: "nope",
	}
	return nil
}
```
