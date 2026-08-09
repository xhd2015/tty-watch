## Expected

- `err` is nil.
- SessionID `full-sess`; Home `/tmp/tty-home-x`; RegistrySubdir `grok-tty-registry`.
- KeepAlive true.
- ExtraPaths `["/opt/bin", "/opt/extra"]`.
- CommandEnv `["AGENT_COLOR=1"]`; CommandUnset `["HTTP_PROXY"]`.
- Command is pure `["codex", "run", "--model", "gpt-5"]` (no serve flags).

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
	p := resp.Parsed
	if p.SessionID != "full-sess" {
		t.Fatalf("SessionID = %q, want full-sess", p.SessionID)
	}
	if p.Home != "/tmp/tty-home-x" {
		t.Fatalf("Home = %q, want /tmp/tty-home-x", p.Home)
	}
	if p.RegistrySubdir != "grok-tty-registry" {
		t.Fatalf("RegistrySubdir = %q, want grok-tty-registry", p.RegistrySubdir)
	}
	if !p.KeepAlive {
		t.Fatal("KeepAlive = false, want true")
	}
	if !equalStrings(p.ExtraPaths, []string{"/opt/bin", "/opt/extra"}) {
		t.Fatalf("ExtraPaths = %#v", p.ExtraPaths)
	}
	if !equalStrings(p.CommandEnv, []string{"AGENT_COLOR=1"}) {
		t.Fatalf("CommandEnv = %#v", p.CommandEnv)
	}
	if !equalStrings(p.CommandUnset, []string{"HTTP_PROXY"}) {
		t.Fatalf("CommandUnset = %#v", p.CommandUnset)
	}
	wantCmd := []string{"codex", "run", "--model", "gpt-5"}
	if !equalStrings(p.Command, wantCmd) {
		t.Fatalf("Command = %#v, want %#v", p.Command, wantCmd)
	}
}
```
