# Scenario

**Feature**: rapid concurrent attach writes are serialized; both complete markers present

```
# two attachers write distinct end-markers nearly simultaneously; no interleaved garbage
harness -> detached cat -> attach A + B concurrent writes -> both *_END markers intact
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-concurrent-ordered"
	req.AttachInput = "CONCURRENT_MARKER_A_END"
	req.AttachInputB = "CONCURRENT_MARKER_B_END"
	return nil
}
```