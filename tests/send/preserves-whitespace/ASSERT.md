## Expected

- Exit code 0.
- Captured bytes are exactly `  spaced  ` (two leading + two trailing spaces).

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
		t.Fatalf("send exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	want := []byte("  spaced  ")
	if !bytes.Equal(resp.InjectedBytes, want) {
		t.Fatalf("injected bytes %q (%v), want %q (%v)", resp.InjectedBytes, resp.InjectedBytes, want, want)
	}
}
```