# Scenario

**Feature**: origin (0,0) maps to CSI 1;1 with default release

```
EncodeSGRClick(0, 0, 0, true) -> \x1b[<0;1;1M\x1b[<0;1;1m
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.EncodeRow = 0
	req.EncodeCol = 0
	req.EncodeBtn = 0
	req.EncodeRelease = true
	return nil
}
```
