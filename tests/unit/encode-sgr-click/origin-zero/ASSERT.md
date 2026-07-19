## Expected

- Bytes exactly `\x1b[<0;1;1M\x1b[<0;1;1m`.

```go
import (
	"bytes"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	want := []byte("\x1b[<0;1;1M\x1b[<0;1;1m")
	if !bytes.Equal(resp.Bytes, want) {
		t.Fatalf("EncodeSGRClick = %q (%v), want %q", resp.Bytes, resp.Bytes, want)
	}
}
```
