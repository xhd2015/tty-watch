# Scenario

**Feature**: DefaultTTYWatchHome joins injected user home with `.tty-watch`

```
UserHome=/Users/alice-fixture -> DefaultTTYWatchHome -> /Users/alice-fixture/.tty-watch
# parallel-safe: no UserHomeDir / getenv
```

## Steps

1. Set a fixture absolute `UserHome` string (not from the process environment).

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.UserHome = "/Users/alice-fixture"
	return nil
}
```
