## Expected

- `BuiltArgv` equals exactly `["min", "--", "sleep", "1"]`.
- No optional serve flag names appear.

## Errors

- None from `Run`.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	if err != nil {
		t.Fatalf("BuildServeArgv: unexpected error: %v", err)
	}
	want := []string{"min", "--", "sleep", "1"}
	if !equalStrings(resp.BuiltArgv, want) {
		t.Fatalf("BuiltArgv = %#v, want %#v", resp.BuiltArgv, want)
	}
	joined := strings.Join(resp.BuiltArgv, "\x00")
	for _, flag := range []string{"--home", "--registry-subdir", "--keep-alive", "--extra-path", "--env", "--unset-env"} {
		if strings.Contains(joined, flag) {
			t.Fatalf("default build must omit %s: %#v", flag, resp.BuiltArgv)
		}
	}
}
```
