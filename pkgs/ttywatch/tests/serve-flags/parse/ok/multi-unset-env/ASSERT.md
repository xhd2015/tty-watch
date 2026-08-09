## Expected

- `err` is nil.
- `Parsed.CommandUnset` equals `["FOO", "BAR"]`.
- `Parsed.Command` is `["true"]`.
- SessionID is `unset-sess`.

## Errors

- None.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	if err != nil {
		t.Fatalf("ParseServeArgv: unexpected error: %v", err)
	}
	if resp.Parsed.SessionID != "unset-sess" {
		t.Fatalf("SessionID = %q, want unset-sess", resp.Parsed.SessionID)
	}
	wantUnset := []string{"FOO", "BAR"}
	if !equalStrings(resp.Parsed.CommandUnset, wantUnset) {
		t.Fatalf("CommandUnset = %#v, want %#v", resp.Parsed.CommandUnset, wantUnset)
	}
	if !equalStrings(resp.Parsed.Command, []string{"true"}) {
		t.Fatalf("Command = %#v, want [true]", resp.Parsed.Command)
	}
}
```
