# Scenario

**Feature**: `--no-release` injects press-only SGR (no lowercase `m` release)

```
tty-watch send <id> --click --row 10 --col 67 --no-release
  -> inject \x1b[<0;68;11M only
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.HasClickRow = true
	req.ClickRow = 10
	req.HasClickCol = true
	req.ClickCol = 67
	req.NoRelease = true
	return nil
}
```
