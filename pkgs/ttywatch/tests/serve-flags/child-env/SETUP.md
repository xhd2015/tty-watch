# Scenario

**Feature**: ServeArgv maps onto ServeOptions for PTY child spawn fields

```
ParseServeArgv(args) -> ServeOptionsFromArgv
  -> ServeOptions.{SessionID,Home,RegistrySubdir,KeepAlive,ExtraPaths,CommandEnv,CommandUnset,Command}
# CommandEnv/Unset feed MergeProcessEnv at spawn (P1 replace; not exercised here)
```

## Steps

1. Set Mode `child-env`.
2. Leaf sets Args that include env/unset and knobs.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "child-env"
	return nil
}
```
