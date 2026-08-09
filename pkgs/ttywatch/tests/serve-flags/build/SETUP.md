# Scenario

**Feature**: BuildServeArgv / BuildServeChildEnv from HeadlessRunOptions

```
HeadlessRunOptions + SessionID
  -> BuildServeArgv -> [sid flags… -- command…]
  -> BuildServeChildEnv(base) -> env without TTY_WATCH_* invent/clear
```

## Steps

1. Default Mode for build leaves is `build`; env-invent leaf overrides to `build-env`.
2. Leaf sets SessionID, Opts, and optionally BaseEnv.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Default surface is argv build; no-tty-watch-env-invent overrides Mode.
	if req.Mode == "" {
		req.Mode = "build"
	}
	return nil
}
```
