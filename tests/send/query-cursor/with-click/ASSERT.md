## Expected

- Product error is **non-nil**.
- Error message indicates exclusive modes (mentions click and/or query).

## Errors

non-nil; keywords: click / query

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
	if !strings.Contains(lower, "click") && !strings.Contains(lower, "query") {
		t.Fatalf("expected click/query exclusive error, got %q", errorText(resp))
	}
}
```
