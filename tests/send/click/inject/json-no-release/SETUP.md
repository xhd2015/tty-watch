# Scenario

**Feature**: `--no-release --json` reports release:false and injects press only

```
tty-watch send <id> --click --row 10 --col 67 --no-release --json
  -> inject press only; stdout release:false
```

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.HasClickRow = true
	req.ClickRow = 10
	req.HasClickCol = true
	req.ClickCol = 67
	req.NoRelease = true
	req.JSON = true
	return nil
}
```
