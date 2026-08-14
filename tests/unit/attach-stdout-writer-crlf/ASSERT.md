## Expected

- `attachStdoutWriter` calls `normalizeTTYOutput` before writing when `rawTTY` is true.
- LF-only screen snapshot text must not pass through unchanged on interactive stdout.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	if !resp.SourceCheckOK {
		t.Fatalf("source check failed: %s", resp.SourceCheckNote)
	}
	if !resp.AttachStdoutWriterNormalizesRawTTY {
		t.Fatalf("attachStdoutWriter must normalize LF-only output on raw TTY: %s", resp.SourceCheckNote)
	}
}
```