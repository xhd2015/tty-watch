# Scenario

**Feature**: screen snapshot text must left-align `[Terminal exited]` when scrollback ends with LF+LF exit marker

```
scrollback: yes\n\n[Terminal exited]
  -> snapshot text lines must be ["yes", "[Terminal exited]"] at column 0
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "unit-screen-snapshot-exit-marker"
	return nil
}
```