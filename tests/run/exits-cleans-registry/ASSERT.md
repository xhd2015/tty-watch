## Expected

- Host exits after child `true` completes.
- Registry has no remaining session entries (`resp.RegistryIDs` empty).
- `resp.RegistryExists` is false.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	if resp.RegistryExists {
		t.Fatalf("registry should be cleaned after attached exit, ids %v", resp.RegistryIDs)
	}
	if len(resp.RegistryIDs) != 0 {
		t.Fatalf("expected empty registry, got %v", resp.RegistryIDs)
	}
}
```