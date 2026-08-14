# Scenario

**Feature**: attach A write is visible to attach B (output broadcast)

```
# two attachers; A writes; B observes A's marker without writing it
harness -> detached cat -> attach A + attach B -> A writes -> B output has marker
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-visible-to-other"
	req.AttachInput = "ATTACH_WRITE_MARKER_A"
	return nil
}
```