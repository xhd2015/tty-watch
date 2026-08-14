# Scenario

**Feature**: send requires session-id and a text message (or a mode flag)

```
cli.Main(["send"] | ["send", sid], …) -> non-nil error
# message: send: requires <session-id> and <message>
```

## Preconditions

- No click/query flags; pure text-mode missing-args path.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "cli"
	// Leaves set Args to ["send"] or ["send", sid] without message/mode flags.
	return nil
}
```
