# Scenario

**Feature**: button field carries btn=2 with default release

```
EncodeSGRClick(10, 67, 2, true) -> \x1b[<2;68;11M\x1b[<2;68;11m
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.EncodeRow = 10
	req.EncodeCol = 67
	req.EncodeBtn = 2
	req.EncodeRelease = true
	return nil
}
```
