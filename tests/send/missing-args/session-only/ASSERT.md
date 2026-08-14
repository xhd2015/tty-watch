## Expected

- Product error is **non-nil**.
- Error message mentions session and message requirement (same family as no-args).

## Errors

non-nil; keywords: send, session, message

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
	text := errorText(resp)
	lower := strings.ToLower(strings.TrimSpace(text))
	if !strings.Contains(lower, "send") {
		t.Fatalf("error should mention send, got %q", text)
	}
	if !strings.Contains(lower, "session") || !strings.Contains(lower, "message") {
		t.Fatalf("error should mention session and message, got %q", text)
	}
}
```
