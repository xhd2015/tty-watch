# Scenario

**Feature**: attach writes while `run` holds screen writer; output visible to run and watch

```
# run sleep attached (screen writer); separate attach writes; run stdout + watch see marker
harness -> run sleep (writer) -> watch + attach -> attach writes -> run + watch output
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "attach-visible-while-run"
	req.RunCommand = []string{"sleep", "300"}
	req.AttachInput = "ATTACH_WRITE_MARKER_A"
	return nil
}
```