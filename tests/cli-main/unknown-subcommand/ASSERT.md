## Expected

- Product error is **non-nil**.
- Error message (prefer `resp.ErrMsg`) mentions `unknown` and/or `subcommand`
  (and/or `usage`).

## Errors

non-nil; keywords: unknown / subcommand / usage

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
	if !strings.Contains(lower, "unknown") &&
		!strings.Contains(lower, "subcommand") &&
		!strings.Contains(lower, "usage") {
		t.Fatalf("expected unknown subcommand error, got %q", errorText(resp))
	}
}
```
