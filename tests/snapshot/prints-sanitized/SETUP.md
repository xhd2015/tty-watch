# Scenario

**Feature**: snapshot prints sanitized text without escape sequences

```
# session prints ANSI red + plain line
harness -> detached printf ANSI -> snapshot -> PLAIN_LINE without escapes
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "snapshot-sanitize"
	return nil
}
```