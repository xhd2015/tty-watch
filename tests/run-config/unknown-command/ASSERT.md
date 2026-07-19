## Expected

- `cli.Run` returns **non-nil** error.
- Error message mentions unknown and/or command/subcommand.

## Errors

non-nil; keywords: unknown / command / subcommand

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
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
