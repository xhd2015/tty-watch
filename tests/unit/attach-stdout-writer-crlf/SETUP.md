# Scenario

**Feature**: attachStdoutWriter applies normalizeTTYOutput on interactive TTY stdout

```
screen snapshot text (LF-only) -> attachStdoutWriter rawTTY path -> CRLF on stdout
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "unit-attach-stdout-writer-crlf"
	return nil
}
```