## Expected

- `cli.Run` returns **non-nil** error.
- Error message mentions **col** (missing/invalid coordinate).

## Errors

non-nil; keyword: col

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
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
