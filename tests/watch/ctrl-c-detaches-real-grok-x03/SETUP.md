# Scenario

**Feature**: Ctrl-C (`\x03`) detaches watch on real grok after alt-screen modes

```
# user flow: tty-watch run grok -> detach -> watch -> Ctrl-C
harness -> run grok -> detach -> watch -> inject \x03 after 1049h+?u seen
```

Documents the user-reported scenario more closely than `ctrl-c-detaches` (which uses
`sleep`) but still **injects** `\x03` via PTY master — not a physical keypress in a
real terminal emulator.

```go
import (
	"os/exec"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if _, err := exec.LookPath("grok"); err != nil {
		t.Skip("grok not found in PATH")
	}
	req.Phase = "watch-ctrl-c-detaches-real-grok-x03"
	return nil
}
```