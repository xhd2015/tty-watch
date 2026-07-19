# Scenario

**Feature**: `--json` prints JSON ack on success while still injecting press+release

```
tty-watch send <id> --click --row 10 --col 67 --json
  -> inject press+release; stdout {"ok":true,"row":10,"col":67,"mouse":0,"release":true}
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.HasClickRow = true
	req.ClickRow = 10
	req.HasClickCol = true
	req.ClickCol = 67
	req.JSON = true
	return nil
}
```
