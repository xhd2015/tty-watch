## Expected

- Exit code 0 (idempotent kill when registry entry is already gone).

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("kill missing expected exit 0, got %d combined %q", resp.ExitCode, resp.Combined)
	}
}
```