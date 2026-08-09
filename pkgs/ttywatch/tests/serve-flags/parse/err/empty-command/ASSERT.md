## Expected

- `err` is non-nil.
- Error message mentions command (e.g. contains `command` case-insensitive).

## Errors

- Parse fails: empty command after `--`.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	_ = resp
	if err == nil {
		t.Fatal("ParseServeArgv: want error for empty command after --, got nil")
	}
	msg := strings.ToLower(err.Error())
	if !strings.Contains(msg, "command") {
		t.Fatalf("error %q should mention command", err.Error())
	}
}
```
