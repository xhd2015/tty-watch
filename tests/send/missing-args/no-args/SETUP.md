# Scenario

**Feature**: `send` with zero trailing args is an error

```
cli.Main(["send"], …) -> error; mentions session + message
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send"}
	return nil
}
```
