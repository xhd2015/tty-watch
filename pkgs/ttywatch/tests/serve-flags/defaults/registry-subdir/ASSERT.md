## Expected

- `DefaultSubdir` is exactly `"registry"`.

## Errors

- None from `Run`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	if err != nil {
		t.Fatalf("defaults: unexpected error: %v", err)
	}
	if resp.DefaultSubdir != "registry" {
		t.Fatalf("DefaultRegistrySubdir() = %q, want %q", resp.DefaultSubdir, "registry")
	}
}
```
