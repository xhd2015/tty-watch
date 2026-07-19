# Scenario

**Feature**: `--query-cursor --json` prints JSON cursor after CUP 5;3

```
tty-watch send <id> --query-cursor --json -> {"row":4,"col":2}
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.JSON = true
	return nil
}
```
