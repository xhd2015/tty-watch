# Scenario

**Feature**: full HeadlessRunOptions emit all flags, `--`, and pure command

```
BuildServeArgv(sid, opts{Home, Subdir, KeepAlive, ExtraPaths, CommandEnv, CommandUnset, Command})
  -> [sid --home … --registry-subdir … --keep-alive
      --extra-path … --env … --unset-env … -- pure-cmd…]
```

## Steps

1. Fill SessionID and Opts with every knob + multi-word command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "build"
	req.SessionID = "build-full"
	req.Opts = ttywatch.HeadlessRunOptions{
		Home:           "/data/home",
		RegistrySubdir: "custom-reg",
		KeepAlive:      true,
		ExtraPaths:     []string{"/x/bin", "/y/bin"},
		CommandEnv:     []string{"FOO=bar", "BAZ=1"},
		CommandUnset:   []string{"NOPE"},
		Command:        []string{"agent", "run", "--verbose"},
	}
	return nil
}
```
