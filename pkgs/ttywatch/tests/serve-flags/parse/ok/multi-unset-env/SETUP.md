# Scenario

**Feature**: repeated `--unset-env KEY` preserved in order

```
Args [sid --unset-env FOO --unset-env BAR -- unset-demo]
  -> CommandUnset [FOO, BAR]
```

## Steps

1. Set two `--unset-env` flags before `--` and command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{
		"unset-sess",
		"--unset-env", "FOO",
		"--unset-env", "BAR",
		"--",
		"true",
	}
	return nil
}
```
