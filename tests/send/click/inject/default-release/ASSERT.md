## Expected

- Exit code 0.
- Stdout and stderr empty (success silent).
- Injected bytes exactly `\x1b[<0;68;11M\x1b[<0;68;11m` (0-based 10,67 → 1-based 11,68).

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
		t.Fatalf("send --click exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	if resp.Stdout != "" || resp.Stderr != "" {
		t.Fatalf("success should be silent; stdout=%q stderr=%q", resp.Stdout, resp.Stderr)
	}
	want := []byte("\x1b[<0;68;11M\x1b[<0;68;11m")
	if !bytes.Equal(resp.InjectedBytes, want) {
		t.Fatalf("injected %q (%v), want %q (%v)", resp.InjectedBytes, resp.InjectedBytes, want, want)
	}
}
```
