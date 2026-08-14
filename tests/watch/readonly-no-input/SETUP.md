# Scenario

**Feature**: watch does not forward stdin to the remote session

```
# cat session would echo stdin; watch must not send harness probe input
harness -> detached cat -> watch + stdin probe -> no SHOULD_NOT_ECHO in output
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "watch-readonly"
	return nil
}
```