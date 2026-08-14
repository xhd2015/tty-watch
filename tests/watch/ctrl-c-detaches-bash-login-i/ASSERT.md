## Expected

- Injected `\x03` detaches watch on bash --login -i session; bash survives.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.TimedOut {
		t.Fatal("watch did not detach on Ctrl-C within 3s for bash --login -i")
	}
	if !resp.RegistryExists || !resp.SessionRunning {
		t.Fatalf("bash session %s not alive after detach", resp.SessionID)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("watch exit code = %d, want 0", resp.ExitCode)
	}
}
```