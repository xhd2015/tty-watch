## Expected

- `err` is nil (bare command without `--` is accepted when no serve flags).
- SessionID `bare-sess`.
- Command `["sh", "-c", "echo hi"]`.
- All knobs empty / KeepAlive false.

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
		t.Fatalf("ParseServeArgv bare command: unexpected error: %v", err)
	}
	if resp.Parsed.SessionID != "bare-sess" {
		t.Fatalf("SessionID = %q, want bare-sess", resp.Parsed.SessionID)
	}
	want := []string{"sh", "-c", "echo hi"}
	if !equalStrings(resp.Parsed.Command, want) {
		t.Fatalf("Command = %#v, want %#v", resp.Parsed.Command, want)
	}
	if resp.Parsed.KeepAlive || resp.Parsed.Home != "" || resp.Parsed.RegistrySubdir != "" {
		t.Fatalf("unexpected knobs set: %+v", resp.Parsed)
	}
}
```
