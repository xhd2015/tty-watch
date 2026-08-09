# Scenario

**Feature**: CommandEnv / CommandUnset and knobs map 1:1 onto ServeOptions

```
Args with --env / --unset-env / --home / --keep-alive / command
  -> ServeOptionsFromArgv(parsed)
  -> fields equal Parsed (spawn-ready struct)
```

## Steps

1. Provide a full-enough argv: home, keep-alive, multi env, unset, pure command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{
		"map-sess",
		"--home", "/map/home",
		"--registry-subdir", "map-reg",
		"--keep-alive",
		"--extra-path", "/map/bin",
		"--env", "COLOR=1",
		"--env", "MODE=test",
		"--unset-env", "PROXY",
		"--",
		"my-agent", "start",
	}
	return nil
}
```
