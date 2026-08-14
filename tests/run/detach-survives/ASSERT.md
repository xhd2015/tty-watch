## Expected

- `resp.RegistryExists` true after detach.
- `resp.SessionRunning` true (listen addr reachable).
- `resp.ListOutput` mentions detached `session-N` id.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	if resp.SessionID == "" {
		t.Fatal("expected detached session id")
	}
	if !resp.RegistryExists {
		t.Fatalf("registry should survive detach for %s", resp.SessionID)
	}
	if !resp.SessionRunning {
		t.Fatalf("session %s should still be reachable after detach", resp.SessionID)
	}
	if !strings.Contains(resp.ListOutput, resp.SessionID) {
		t.Fatalf("list should include %s, got %q", resp.SessionID, resp.ListOutput)
	}
}
```