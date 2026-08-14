# Scenario

**Feature**: watch streams raw PTY output to stdout

```
# detached loop echoes marker
harness -> detached echo loop -> watch -> WATCH_MARKER on stdout
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "watch-stream"
	req.WatchProbe = "4s"
	return nil
}
```