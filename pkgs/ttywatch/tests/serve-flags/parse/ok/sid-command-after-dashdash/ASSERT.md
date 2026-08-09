## Expected

- `err` is nil.
- `Parsed.SessionID` is `sess-1`.
- `Parsed.Command` is `["echo", "hello"]`.
- Home, RegistrySubdir empty; KeepAlive false; ExtraPaths / CommandEnv / CommandUnset empty.

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
	if resp.Parsed.SessionID != "sess-1" {
		t.Fatalf("SessionID = %q, want sess-1", resp.Parsed.SessionID)
	}
	if !equalStrings(resp.Parsed.Command, []string{"echo", "hello"}) {
		t.Fatalf("Command = %#v, want [echo hello]", resp.Parsed.Command)
	}
	if resp.Parsed.Home != "" {
		t.Fatalf("Home = %q, want empty when --home omitted", resp.Parsed.Home)
	}
	if resp.Parsed.RegistrySubdir != "" {
		t.Fatalf("RegistrySubdir = %q, want empty when flag omitted", resp.Parsed.RegistrySubdir)
	}
	if resp.Parsed.KeepAlive {
		t.Fatal("KeepAlive = true, want false when --keep-alive omitted")
	}
	if len(resp.Parsed.ExtraPaths) != 0 {
		t.Fatalf("ExtraPaths = %#v, want empty", resp.Parsed.ExtraPaths)
	}
	if len(resp.Parsed.CommandEnv) != 0 {
		t.Fatalf("CommandEnv = %#v, want empty", resp.Parsed.CommandEnv)
	}
	if len(resp.Parsed.CommandUnset) != 0 {
		t.Fatalf("CommandUnset = %#v, want empty", resp.Parsed.CommandUnset)
	}
}
```
