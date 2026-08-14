# Scenario

**Feature**: free-text args cannot mix with `--click`

```
cli.Main(["send", "sess-flags", "--click", "--row", "1", "--col", "1", "hello"], …)
  -> error; cannot mix free-text with click
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags", "--click", "--row", "1", "--col", "1", "hello"}
	return nil
}
```
