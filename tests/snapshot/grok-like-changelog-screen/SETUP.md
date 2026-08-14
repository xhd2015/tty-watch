# Scenario

**Bug**: `tty-watch snapshot` on grok-style changelog alt-screen TUIs prefers garbled
scrollback replay over the ptywrap screen frame (menu shows `Quit q` without `ctrl+q`,
missing bordered changelog box, prompt, and footer).

```
# grok-like alt-screen changelog with absolute CUP redraws
harness -> run --detach (grok-like sh) -> snapshot -> plain text matches attach screen
```

Reproduces user report: detached `grok` session snapshot shows changelog-only smear
instead of the full bordered UI that `attach`/`watch` mirror.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "snapshot-grok-like-changelog-screen"
	return nil
}
```