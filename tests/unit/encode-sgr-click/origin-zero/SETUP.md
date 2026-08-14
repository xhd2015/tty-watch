# Scenario

**Feature**: origin (0,0) maps to CSI 1;1 with default release

```
EncodeSGRClick(0, 0, 0, true) -> \x1b[<0;1;1M\x1b[<0;1;1m
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.EncodeRow = 0
	req.EncodeCol = 0
	req.EncodeBtn = 0
	req.EncodeRelease = true
	return nil
}
```
