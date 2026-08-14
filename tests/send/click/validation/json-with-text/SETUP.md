# Scenario

**Feature**: `--json` is rejected with text-only send mode

```
cli.Main(["send", "sess-flags", "--json", "hello"], …)
  -> error; --json only for click/query-cursor
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"send", "sess-flags", "--json", "hello"}
	return nil
}
```
