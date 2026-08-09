# Scenario

**Feature**: successful ParseServeArgv paths (non-empty command)

```
req.Args with valid sid + flags + command -> ParseServeArgv -> ok ServeArgv
# Assert: err==nil and Parsed fields
```

## Steps

1. Expect `Run` error to be nil.
2. Leaf sets concrete `req.Args` fixture.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Mode already "parse" from parent; ok branch expects success.
	if req.Mode != "parse" {
		t.Fatalf("parse/ok expects Mode=parse, got %q", req.Mode)
	}
	return nil
}
```
