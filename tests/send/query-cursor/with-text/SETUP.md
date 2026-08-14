# Scenario

**Feature**: free-text cannot mix with `--query-cursor`

```
cli.Main(["send", "sess-flags", "--query-cursor", "hello"], …)
  -> error; cannot mix free-text with query-cursor
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "cli"
	req.Args = []string{"send", "sess-flags", "--query-cursor", "hello"}
	return nil
}
```
