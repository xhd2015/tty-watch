# Scenario

**Feature**: `--click --col` without `--row` is an error

```
cli.Main(["send", "sess-flags", "--click", "--col", "1"], …) -> error; mentions row
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags", "--click", "--col", "1"}
	return nil
}
```
