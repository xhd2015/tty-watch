## Expected

- Product error is **non-nil**.
- Error message mentions row and/or negative/invalid.

## Errors

non-nil; keywords: row / negative / invalid

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
	if !strings.Contains(lower, "row") &&
		!strings.Contains(lower, "negative") &&
		!strings.Contains(lower, "invalid") {
		t.Fatalf("expected negative/invalid row error, got %q", errorText(resp))
	}
}
```
