# Scenario

**Feature**: attach forwards typed stdin to the PTY master

```
# detached cat; attach types unique marker; marker echoed on attach stdout
harness -> detached cat -> tty-watch attach <id> + stdin -> ATTACH_STDIN_MARKER visible
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-forwards-stdin"
	req.AttachInput = "ATTACH_STDIN_MARKER"
	return nil
}
```