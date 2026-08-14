# Scenario

**Feature**: attach write is broadcast to readonly watch observers

```
# watch connected; attach writes marker; watch stdout contains marker
harness -> detached cat -> watch + attach -> attach writes -> watch sees marker
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-visible-to-watch"
	req.AttachInput = "ATTACH_WRITE_MARKER_A"
	return nil
}
```