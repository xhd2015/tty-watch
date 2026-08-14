# Scenario

**Bug**: normalizeTTYOutput must expand bare LF to CRLF while preserving standalone CR

```
LF-only screen snapshot text -> normalizeTTYOutput -> CRLF on raw TTY stdout
standalone \r (in-place redraw) -> unchanged; trailing bare \n -> \r\n
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "unit-normalize-tty-output"
	return nil
}
```