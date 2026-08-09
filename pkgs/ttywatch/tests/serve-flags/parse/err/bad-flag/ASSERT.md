## Expected

- `err` is non-nil.
- Error surfaces the bad flag name or generic unknown-flag wording.

## Errors

- Parse fails on `--not-a-real-flag`.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	_ = resp
	if err == nil {
		t.Fatal("ParseServeArgv: want error for unknown flag, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "not-a-real-flag") &&
		!strings.Contains(strings.ToLower(msg), "unknown") &&
		!strings.Contains(strings.ToLower(msg), "flag") {
		t.Fatalf("error %q should mention the bad flag or that a flag is unknown", msg)
	}
}
```
