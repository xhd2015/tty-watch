# Scenario

**Feature**: `--mouse 2` puts button 2 in both press and release SGR

```
tty-watch send <id> --click --row 10 --col 67 --mouse 2
  -> inject \x1b[<2;68;11M\x1b[<2;68;11m
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
	req.HasMouse = true
	req.Mouse = 2
	return nil
}
```
