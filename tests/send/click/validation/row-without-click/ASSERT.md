## Expected

- Product error is **non-nil**.
- Error message indicates `--row`/`--col` require `--click` (must contain "click" so
  dummy sid `sess-flags` cannot false-pass via the word "flag").

## Errors

non-nil; keyword: click

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
	if !strings.Contains(lower, "click") {
		t.Fatalf("expected error that --row/--col require --click, got %q", errorText(resp))
	}
}
```
