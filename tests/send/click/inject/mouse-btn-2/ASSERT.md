## Expected

- Exit code 0; silent stdout.
- Injected bytes use button field `2` in press and release.

```go
import (
	"bytes"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("send --click --mouse 2 exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	if resp.Stdout != "" || resp.Stderr != "" {
		t.Fatalf("success should be silent; stdout=%q stderr=%q", resp.Stdout, resp.Stderr)
	}
	want := []byte("\x1b[<2;68;11M\x1b[<2;68;11m")
	if !bytes.Equal(resp.InjectedBytes, want) {
		t.Fatalf("injected %q (%v), want %q (%v)", resp.InjectedBytes, resp.InjectedBytes, want, want)
	}
}
```
