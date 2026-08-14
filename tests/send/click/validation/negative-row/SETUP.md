# Scenario

**Feature**: negative `--row` is rejected

```
cli.Main(["send", "sess-flags", "--click", "--row", "-1", "--col", "1"], …) -> error
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags", "--click", "--row", "-1", "--col", "1"}
	return nil
}
```
