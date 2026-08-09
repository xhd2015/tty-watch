# Scenario

**Feature**: BuildServeChildEnv never invents or clears the four re-exec knobs

```
base [PATH=/bin, TTY_WATCH_HOME=/ambient]
opts {Home, RegistrySubdir, KeepAlive, ExtraPaths}  # would have invented under old policy
  -> BuildServeChildEnv
  -> does not invent TTY_WATCH_* from opts
  -> does not clear ambient TTY_WATCH_HOME
```

## Steps

1. Mode `build-env`; BaseEnv includes ambient `TTY_WATCH_HOME` and a normal `PATH`.
2. Opts set all four knobs so old `serveChildEnv` would invent/clear.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "build-env"
	req.BaseEnv = []string{
		"PATH=/bin",
		"TTY_WATCH_HOME=/ambient-home",
	}
	req.Opts = ttywatch.HeadlessRunOptions{
		Home:           "/from-opts-home",
		RegistrySubdir: "from-opts-sub",
		KeepAlive:      true,
		ExtraPaths:     []string{"/from/opts/path"},
		Command:        []string{"true"},
	}
	return nil
}
```
