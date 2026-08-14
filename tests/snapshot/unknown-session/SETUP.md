# Scenario

**Feature**: snapshot on unknown session id fails

```
# missing registry entry
tty-watch snapshot session-99999 -> error
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "snapshot-missing"
	req.SnapshotID = "session-99999"
	return nil
}
```