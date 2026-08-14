# Scenario

**Feature**: attach resize updates session PTY dimensions (last-wins)

```
# wide-line session; attach resize cols=100; snapshot shows single unwrapped line
harness -> run wide-line sh -> detach -> attach resize -> snapshot not 80-col wrapped
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-resize"
	return nil
}
```