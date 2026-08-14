## Expected

- Exit code 0.
- List output contains no `session-` id lines.

```go
import (
	"regexp"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("list empty exit %d, output %q", resp.ExitCode, resp.ListOutput)
	}
	if regexp.MustCompile(`session-\d+`).MatchString(resp.ListOutput) {
		t.Fatalf("expected no sessions, got %q", resp.ListOutput)
	}
}
```