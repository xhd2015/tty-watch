## Expected

- Exit code 0; silent stdout/stderr.
- Injected bytes exactly `\x1b[<0;68;11M` (no release).

```go
import (
	"bytes"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("send --click --no-release exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	if resp.Stdout != "" || resp.Stderr != "" {
		t.Fatalf("success should be silent; stdout=%q stderr=%q", resp.Stdout, resp.Stderr)
	}
	want := []byte("\x1b[<0;68;11M")
	if !bytes.Equal(resp.InjectedBytes, want) {
		t.Fatalf("injected %q (%v), want %q (%v)", resp.InjectedBytes, resp.InjectedBytes, want, want)
	}
}
```
