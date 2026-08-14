## Expected

- Bytes exactly `\x1b[<0;1;1M\x1b[<0;1;1m`.

```go
import (
	"bytes"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("\x1b[<0;1;1M\x1b[<0;1;1m")
	if !bytes.Equal(resp.Bytes, want) {
		t.Fatalf("EncodeSGRClick = %q (%v), want %q", resp.Bytes, resp.Bytes, want)
	}
}
```
