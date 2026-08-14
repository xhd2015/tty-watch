## Expected

- `ParseArgs` returns **nil** error.
- `ParsedConfig.Command` is `"send"`.
- `ParsedConfig.Send` is non-nil with:
  - `Session == "sid"`
  - `Mode == cli.SendModeClick`
  - `Row == 10`, `Col == 67`
  - `NoRelease == false` (default release ON)
  - `Mouse == 0` (default button)

## Errors

nil

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"

	"github.com/xhd2015/tty-watch/cli"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	assertOK(t, resp)
	cfg := resp.ParsedConfig
	if cfg.Command != "send" {
		t.Fatalf("Command = %q, want send", cfg.Command)
	}
	if cfg.Send == nil {
		t.Fatal("ParsedConfig.Send is nil")
	}
	s := cfg.Send
	if s.Session != "sid" {
		t.Fatalf("Send.Session = %q, want sid", s.Session)
	}
	if s.Mode != cli.SendModeClick {
		t.Fatalf("Send.Mode = %v, want SendModeClick", s.Mode)
	}
	if s.Row != 10 {
		t.Fatalf("Send.Row = %d, want 10", s.Row)
	}
	if s.Col != 67 {
		t.Fatalf("Send.Col = %d, want 67", s.Col)
	}
	if s.NoRelease {
		t.Fatal("Send.NoRelease = true, want false (default release ON)")
	}
	if s.Mouse != 0 {
		t.Fatalf("Send.Mouse = %d, want 0", s.Mouse)
	}
}
```
