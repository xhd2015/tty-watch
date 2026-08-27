# Scenario

**Feature**: hard-killed `__serve__` must not leave an HUP-immune PTY child

Crime scene (Marcus / long `tty-watch` / `agent-run` use): after the session
manager dies, `bash`/tool children that ignore `SIGHUP` remain as `PPID=1`
holding `/dev/ttys*` with no live master and exhaust `kern.tty.ptmx_max`.

```
# leak path (current)
tty-watch run --detach --session-id sigkill-no-orphan -- <hup-immune sticky>
  -> SIGKILL __serve__
  -> sticky child still alive (PPID=1, slave FD open)

# expected
same recipe -> CommandAlive == false (child reaped via death-signal / pgid)
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "kill-sigkill-serve-no-orphan-child"
	req.CustomSessionID = "sigkill-no-orphan"
	return nil
}
```
