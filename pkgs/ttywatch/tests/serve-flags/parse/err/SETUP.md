# Scenario

**Feature**: ParseServeArgv error paths (empty command, bad flag)

```
req.Args invalid -> ParseServeArgv -> non-nil error
# Assert: err != nil; no successful command payload required
```

## Steps

1. Expect `Run` to return a non-nil error.
2. Leaf sets the failing argv fixture.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	if req.Mode != "parse" {
		t.Fatalf("parse/err expects Mode=parse, got %q", req.Mode)
	}
	return nil
}
```
