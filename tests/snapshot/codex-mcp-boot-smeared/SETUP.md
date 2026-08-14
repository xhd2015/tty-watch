# Scenario

**Bug**: snapshot taken while codex MCP status line is incrementally redrawn smears
stacked boot progress and leaks kitty CSI fragments.

```
# snapshot mid MCP boot (before final error settles)
harness -> run --session-id codex (codex-cursor-drawn sh) -> detach -> snapshot early
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Phase = "snapshot-codex-mcp-boot-smeared"
	return nil
}
```