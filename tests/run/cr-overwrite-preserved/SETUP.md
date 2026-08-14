# Scenario

**Feature**: carriage-return overwrite renders correctly on writer attach

```
harness PTY -> tty-watch run sh -c "printf 'MARKER_A\\rMARKER_B\\n'" -> MARKER_B visible, not MARKER_AMARKER_B
```

PTY programs and interactive shells rely on `\r` for cursor positioning; stripping
`\r` smears prompts and error lines across the terminal width.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "run-cr-overwrite"
	return nil
}
```