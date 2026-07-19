## Expected

- Product error is **nil**.
- Stdout mentions **Usage** and at least one known subcommand (`send` or `run`).

## Errors

nil

```go
import (
	"strings"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	assertOK(t, resp)
	lower := strings.ToLower(resp.Stdout)
	if lower == "" {
		lower = strings.ToLower(resp.Combined)
	}
	if !strings.Contains(lower, "usage") {
		t.Fatalf("empty-args should print Usage, got %q", resp.Stdout)
	}
	if !strings.Contains(lower, "send") && !strings.Contains(lower, "run") {
		t.Fatalf("empty-args usage should list send or run, got %q", resp.Stdout)
	}
}
```
