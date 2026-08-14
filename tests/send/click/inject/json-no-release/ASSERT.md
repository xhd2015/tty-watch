## Expected

- Exit code 0.
- JSON ack has `release:false`, same row/col/mouse.
- Injected bytes press-only.

```go
import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("send --click --no-release --json exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	var ack struct {
		OK      bool `json:"ok"`
		Row     int  `json:"row"`
		Col     int  `json:"col"`
		Mouse   int  `json:"mouse"`
		Release bool `json:"release"`
	}
	if err := json.Unmarshal([]byte(resp.Stdout), &ack); err != nil {
		t.Fatalf("stdout JSON: %v; stdout=%q", err, resp.Stdout)
	}
	if !ack.OK || ack.Row != 10 || ack.Col != 67 || ack.Mouse != 0 || ack.Release {
		t.Fatalf("ack %+v, want ok row=10 col=67 mouse=0 release=false", ack)
	}
	want := []byte("\x1b[<0;68;11M")
	if !bytes.Equal(resp.InjectedBytes, want) {
		t.Fatalf("injected %q (%v), want %q", resp.InjectedBytes, resp.InjectedBytes, want)
	}
}
```
