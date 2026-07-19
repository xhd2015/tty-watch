# Scenario

**Feature**: default click (row=10,col=67) injects press+release SGR for button 0

```
tty-watch send <id> --click --row 10 --col 67
  -> inject \x1b[<0;68;11M\x1b[<0;68;11m ; silent stdout
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.HasClickRow = true
	req.ClickRow = 10
	req.HasClickCol = true
	req.ClickCol = 67
	// default: release ON, mouse 0 omitted
	return nil
}
```
