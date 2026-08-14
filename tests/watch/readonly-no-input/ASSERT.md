## Expected

- Watch output does not contain `SHOULD_NOT_ECHO` (stdin probe was not forwarded).
- Observer mode is readonly.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	combined := resp.Combined
	if combined == "" {
		combined = resp.Stdout
	}
	if strings.Contains(combined, "SHOULD_NOT_ECHO") {
		t.Fatalf("watch forwarded stdin to session, output %q", combined)
	}
}
```