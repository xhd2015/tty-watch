# Scenario

**Feature**: pure `EncodeSGRClick(row,col,btn,release)` builds SGR mouse wire bytes

```
# sealed API: pkgs/ttywatch.EncodeSGRClick
EncodeSGRClick(row, col, btn, release) -> ESC[<btn;col+1;row+1M [+ m if release]
```

## Preconditions

- No live session; Mode=`encode`.
- 0-based row/col → 1-based in CSI params.
- Import path: `github.com/xhd2015/tty-watch/pkgs/ttywatch`.

## Steps

1. Leaf sets EncodeRow / EncodeCol / EncodeBtn / EncodeRelease.
2. Run returns encoder output in Response.Bytes.
3. Assert compares exact bytes.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "encode"
	return nil
}
```
