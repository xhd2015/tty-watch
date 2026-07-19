## Expected

- Product error is **non-nil**.
- Error message indicates mix/text/message/query conflict.

## Errors

non-nil; keywords: mix / text / message / query

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
	if !strings.Contains(lower, "mix") &&
		!strings.Contains(lower, "text") &&
		!strings.Contains(lower, "message") &&
		!strings.Contains(lower, "query") {
		t.Fatalf("expected mix/text/query conflict, got %q", errorText(resp))
	}
}
```
