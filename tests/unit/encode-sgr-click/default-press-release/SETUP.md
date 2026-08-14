# Scenario

**Feature**: default press+release at row=10 col=67 btn=0

```
EncodeSGRClick(10, 67, 0, true) -> \x1b[<0;68;11M\x1b[<0;68;11m
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
	req.EncodeRelease = true
	return nil
}
```
