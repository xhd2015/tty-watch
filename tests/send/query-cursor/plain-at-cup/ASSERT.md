## Expected

- Exit code 0.
- Stdout exactly `row=4 col=2` (optional trailing newline).
- No inject of CSI 6n into child (not asserted via capture; host-only read).

## Expected Output

```text
row=4 col=2
```

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"

	"github.com/xhd2015/doctest/assert"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("query-cursor exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	out := strings.TrimSpace(resp.Stdout)
	assert.Output(t, out+"\n", `---
version: 3
---
row=4 col=2
`)
}
```
