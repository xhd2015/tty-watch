# Scenario

**Feature**: empty args print root usage like help (nil error)

```
cli.Main([], …) -> nil error; stdout usage (same family as -h)
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{}
	return nil
}
```
