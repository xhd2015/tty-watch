# Scenario

**Feature**: empty knobs omit optional flags — only sid, `--`, command

```
BuildServeArgv("min", opts{Command: [sleep 1]})
  -> [min -- sleep 1]
# no --home / --registry-subdir / --keep-alive / --extra-path / --env / --unset-env
```

## Steps

1. Set SessionID and Command only; leave Home/Subdir empty, KeepAlive false.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "build"
	req.SessionID = "min"
	req.Opts = ttywatch.HeadlessRunOptions{
		Command: []string{"sleep", "1"},
	}
	return nil
}
```
