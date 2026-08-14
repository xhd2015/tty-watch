# Scenario

**Feature**: `--row` / `--col` without `--click` is an error

```
cli.Main(["send", "sess-flags", "--row", "1", "--col", "1"], …)
  -> error; --row/--col require --click
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags", "--row", "1", "--col", "1"}
	return nil
}
```
