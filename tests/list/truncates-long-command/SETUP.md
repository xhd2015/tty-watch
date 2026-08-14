# Scenario

**Feature**: list truncates COMMAND column to 64 characters with ellipsis

```
# detached run with argv joined length > 64
harness -> detached bash -c long echo -> tty-watch list -> COMMAND ends with ...
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "list-fields"
	req.RunCommand = []string{
		"bash", "-c",
		"echo this is a very long command that exceeds sixty four characters total",
	}
	return nil
}
```