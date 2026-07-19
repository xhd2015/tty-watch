## Expected

- Bytes exactly `\x1b[<0;68;11M\x1b[<0;68;11m`
  (0-based (10,67) → wire col=68 row=11; press then release).

```go
import (
	"bytes"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("\x1b[<0;68;11M\x1b[<0;68;11m")
	if !bytes.Equal(resp.Bytes, want) {
		t.Fatalf("EncodeSGRClick = %q (%v), want %q", resp.Bytes, resp.Bytes, want)
	}
}
```
