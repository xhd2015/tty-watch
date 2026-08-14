## Expected

- Product error is **non-nil**.
- Error message mentions json and/or text/message mode incompatibility.

## Errors

non-nil; keywords: json / text / message

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
	if !strings.Contains(lower, "json") &&
		!strings.Contains(lower, "text") &&
		!strings.Contains(lower, "message") {
		t.Fatalf("expected --json with text-mode error, got %q", errorText(resp))
	}
}
```
