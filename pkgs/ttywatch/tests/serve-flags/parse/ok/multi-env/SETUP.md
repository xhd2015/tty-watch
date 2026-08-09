# Scenario

**Feature**: repeated `--env KEY=VALUE` preserved in order

```
Args [sid --env A=1 --env B=2 --env A=3 -- echo]
  -> CommandEnv [A=1, B=2, A=3] (order preserved; merge last-wins is spawn-time)
```

## Steps

1. Set three `--env` flags (including a repeated key) before `--` and command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{
		"env-sess",
		"--env", "A=1",
		"--env", "B=2",
		"--env", "A=3",
		"--",
		"env",
	}
	return nil
}
```
