## Expected

- Exit code 0.
- Registry file removed for killed session.
- Session no longer TCP-reachable.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("kill exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	if resp.RegistryExists {
		t.Fatalf("registry should be removed after kill for %s", resp.SessionID)
	}
	if resp.SessionRunning {
		t.Fatalf("session %s should not be reachable after kill", resp.SessionID)
	}
}
```