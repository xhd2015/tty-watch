## Expected

- Combined PTY/host capture contains no line starting with `session-`.
- Host must not print session id on stdout or stderr during attach or Ctrl-] detach.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	combined := resp.Combined
	if combined == "" {
		combined = resp.Stdout
	}
	assertNoHostSessionID(t, combined)
}
```