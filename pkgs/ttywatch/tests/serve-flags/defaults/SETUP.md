# Scenario

**Feature**: pure default helpers — no ambient TTY_WATCH_* reads

```
DefaultTTYWatchHome(userHome) -> filepath.Join(userHome, ".tty-watch")
DefaultRegistrySubdir() -> "registry"
# inject UserHome; never t.Setenv / getenv of knobs
```

## Steps

1. Set Mode `defaults`.
2. Leaf sets `UserHome` when testing home join.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Mode = "defaults"
	return nil
}
```
