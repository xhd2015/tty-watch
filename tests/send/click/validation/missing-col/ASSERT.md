## Expected

- Product error is **non-nil**.
- Error message mentions **col** (missing/required).

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
		t.Fatalf("expected error about col, got %q", errorText(resp))
	}
}
```
