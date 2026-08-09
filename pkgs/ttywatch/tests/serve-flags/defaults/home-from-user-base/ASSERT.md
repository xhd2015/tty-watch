## Expected

- `DefaultHome` equals `filepath.Join(UserHome, ".tty-watch")` → `/Users/alice-fixture/.tty-watch`.
- Helper is pure: result determined only by the injected base (policy B: no ambient read).

## Errors

- None from `Run`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatalf("defaults: unexpected error: %v", err)
	}
	want := joinHome(req.UserHome)
	if resp.DefaultHome != want {
		t.Fatalf("DefaultTTYWatchHome(%q) = %q, want %q", req.UserHome, resp.DefaultHome, want)
	}
	if resp.DefaultHome == "" || resp.DefaultHome == ".tty-watch" {
		t.Fatalf("DefaultHome must be absolute join, got %q", resp.DefaultHome)
	}
}
```
