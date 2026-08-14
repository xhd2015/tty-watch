# Scenario

**Feature**: release=false emits press only

```
EncodeSGRClick(10, 67, 0, false) -> \x1b[<0;68;11M
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.EncodeRow = 10
	req.EncodeCol = 67
	req.EncodeBtn = 0
	req.EncodeRelease = false
	return nil
}
```
