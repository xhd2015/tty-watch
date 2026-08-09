# Scenario

**Feature**: ParseServeArgv — argv after serve token → ServeArgv

```
# ok and err subtrees share Mode=parse
req.Args -> ParseServeArgv -> ServeArgv or error
```

## Steps

1. Set `req.Mode` to `parse`.
2. Leaf fills `req.Args` (sid-first wire, no serve token).

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "parse"
	return nil
}
```
