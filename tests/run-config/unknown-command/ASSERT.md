## Expected

- `cli.Run` returns **non-nil** error.
- Error message mentions unknown and/or command/subcommand.

## Errors

non-nil; keywords: unknown / command / subcommand

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
	assertErr(t, resp)
	lower := strings.ToLower(errorText(resp))
	if !strings.Contains(lower, "unknown") &&
		!strings.Contains(lower, "command") &&
		!strings.Contains(lower, "subcommand") {
		t.Fatalf("expected unknown command error, got %q", errorText(resp))
	}
}
```
