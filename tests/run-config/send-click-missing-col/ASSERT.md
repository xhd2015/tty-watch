## Expected

- `cli.Run` returns **non-nil** error.
- Error message mentions **col** (missing/invalid coordinate).

## Errors

non-nil; keyword: col

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	assertErr(t, resp)
	lower := strings.ToLower(errorText(resp))
	if !strings.Contains(lower, "col") {
		t.Fatalf("expected Run click missing/invalid col error, got %q", errorText(resp))
	}
}
```
