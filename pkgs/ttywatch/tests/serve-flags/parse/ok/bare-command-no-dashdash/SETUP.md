# Scenario

**Feature**: bare command without `--` (no serve flags after sid)

```
Args [bare-sess sh -c echo hi]  # no -- separator, no --home/…
  -> SessionID=bare-sess, Command=[sh -c echo hi]
# command args may start with "-" (e.g. -c); only known serve flags are parsed
```

## Steps

1. Set sid plus a multi-word agent command with no serve flags and no `--`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Migration path matching pre-P2 re-exec: sid then pure command (may include -c).
	req.Args = []string{"bare-sess", "sh", "-c", "echo hi"}
	return nil
}
```
