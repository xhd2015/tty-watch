## Expected

- Product error is **non-nil**.
- Error message mentions **row**.

## Errors

non-nil; keyword: row

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
	if !strings.Contains(lower, "row") {
		t.Fatalf("expected error about row, got %q", errorText(resp))
	}
}
```
