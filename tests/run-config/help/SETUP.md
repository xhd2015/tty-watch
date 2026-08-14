# Scenario

**Feature**: `Run(Config{Command:"help"})` prints usage on stdout with nil error

```
# empty Command is the same family; this leaf uses Command:"help"
cli.Run(Config{Command: "help", Stdout: buf, …}) -> nil
# stdout: Usage + subcommands
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
		Command: "help",
	}
	return nil
}
```
