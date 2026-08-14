# Scenario

**Feature**: unknown subcommand returns non-nil error

```
cli.Main(["nope"], …) -> error; message mentions unknown / subcommand
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"nope"}
	return nil
}
```
