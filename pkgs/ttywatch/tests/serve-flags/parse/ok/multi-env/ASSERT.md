## Expected

- `err` is nil.
- `Parsed.CommandEnv` equals `["A=1", "B=2", "A=3"]` (all occurrences, order kept).
- `Parsed.Command` is `["env"]`.
- SessionID is `env-sess`.

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
	if resp.Parsed.SessionID != "env-sess" {
		t.Fatalf("SessionID = %q, want env-sess", resp.Parsed.SessionID)
	}
	wantEnv := []string{"A=1", "B=2", "A=3"}
	if !equalStrings(resp.Parsed.CommandEnv, wantEnv) {
		t.Fatalf("CommandEnv = %#v, want %#v", resp.Parsed.CommandEnv, wantEnv)
	}
	if !equalStrings(resp.Parsed.Command, []string{"env"}) {
		t.Fatalf("Command = %#v, want [env]", resp.Parsed.Command)
	}
}
```
