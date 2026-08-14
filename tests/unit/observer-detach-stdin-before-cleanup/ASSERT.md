## Expected

- `streamObserver` detach path must restore stdin termios (`term.Restore`) **before**
  writing observer TTY cleanup sequences to stdout.
- Relying only on `defer term.Restore` after `writeObserverTTYDetachCleanup` is incorrect:
  cleanup must run in cooked stdin mode so iTerm2 consumes CSI instead of echoing `^[[?0u`.

```go
import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	if !resp.SourceCheckOK {
		t.Fatalf("could not analyze attach.go: %s", resp.SourceCheckNote)
	}
	if !resp.StdinRestoredBeforeCleanup {
		t.Fatalf("attach.go writes observer TTY cleanup before stdin term.Restore on detach; causes ^[[?0u prompt smear and kitty typing garbage in iTerm2")
	}
}
```