# Scenario

**Feature**: attached command exit removes registry entry

```
# short-lived command exits while attached
harness PTY -> tty-watch true -> host exits -> registry empty
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-exit-clean"
	req.RunCommand = []string{"true"}
	return nil
}
```