## Expected

- Product error is **non-nil**.
- Error message indicates cannot mix free-text with click (mix / text / message / click).

## Errors

non-nil; keywords: mix / text / message / click

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
		!strings.Contains(lower, "click") {
		t.Fatalf("expected mix/text/click conflict error, got %q", errorText(resp))
	}
}
```
