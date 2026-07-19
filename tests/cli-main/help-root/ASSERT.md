## Expected

- Product error is **nil** (`cli.Main` success).
- Stdout (or combined) mentions **Usage** (case-insensitive) and both subcommands
  **send** and **run**.

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
		t.Fatalf("help stdout should mention Usage, got %q", resp.Stdout)
	}
	if !strings.Contains(lower, "send") {
		t.Fatalf("help stdout should mention send, got %q", resp.Stdout)
	}
	if !strings.Contains(lower, "run") {
		t.Fatalf("help stdout should mention run, got %q", resp.Stdout)
	}
}
```
