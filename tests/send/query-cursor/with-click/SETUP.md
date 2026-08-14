# Scenario

**Feature**: `--click` and `--query-cursor` cannot combine

```
cli.Main(["send", "sess-flags", "--click", "--row", "1", "--col", "1", "--query-cursor"], …)
  -> error; exclusive modes
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "cli"
	req.Args = []string{
		"send", "sess-flags",
		"--click", "--row", "1", "--col", "1",
		"--query-cursor",
	}
	return nil
}
```
