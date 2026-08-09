## Expected

- `err` is nil.
- `ServeOpts` mirrors `Parsed` for SessionID, Home, RegistrySubdir, KeepAlive,
  ExtraPaths, CommandEnv, CommandUnset, Command.
- CommandEnv/CommandUnset are non-empty (fields exist for spawn MergeProcessEnv).

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
		t.Fatalf("child-env map: unexpected error: %v", err)
	}
	p := resp.Parsed
	o := resp.ServeOpts

	if o.SessionID != p.SessionID || o.SessionID != "map-sess" {
		t.Fatalf("ServeOpts.SessionID = %q, Parsed = %q", o.SessionID, p.SessionID)
	}
	if o.Home != p.Home || o.Home != "/map/home" {
		t.Fatalf("ServeOpts.Home = %q, Parsed.Home = %q", o.Home, p.Home)
	}
	if o.RegistrySubdir != p.RegistrySubdir || o.RegistrySubdir != "map-reg" {
		t.Fatalf("ServeOpts.RegistrySubdir = %q, Parsed = %q", o.RegistrySubdir, p.RegistrySubdir)
	}
	if o.KeepAlive != p.KeepAlive || !o.KeepAlive {
		t.Fatalf("ServeOpts.KeepAlive = %v, Parsed = %v", o.KeepAlive, p.KeepAlive)
	}
	if !equalStrings(o.ExtraPaths, p.ExtraPaths) || !equalStrings(o.ExtraPaths, []string{"/map/bin"}) {
		t.Fatalf("ServeOpts.ExtraPaths = %#v, Parsed = %#v", o.ExtraPaths, p.ExtraPaths)
	}
	wantEnv := []string{"COLOR=1", "MODE=test"}
	if !equalStrings(o.CommandEnv, p.CommandEnv) || !equalStrings(o.CommandEnv, wantEnv) {
		t.Fatalf("ServeOpts.CommandEnv = %#v, Parsed = %#v", o.CommandEnv, p.CommandEnv)
	}
	wantUnset := []string{"PROXY"}
	if !equalStrings(o.CommandUnset, p.CommandUnset) || !equalStrings(o.CommandUnset, wantUnset) {
		t.Fatalf("ServeOpts.CommandUnset = %#v, Parsed = %#v", o.CommandUnset, p.CommandUnset)
	}
	wantCmd := []string{"my-agent", "start"}
	if !equalStrings(o.Command, p.Command) || !equalStrings(o.Command, wantCmd) {
		t.Fatalf("ServeOpts.Command = %#v, Parsed = %#v", o.Command, p.Command)
	}
}
```
