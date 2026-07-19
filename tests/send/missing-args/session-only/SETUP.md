# Scenario

**Feature**: `send` with session id but no message/mode is an error

```
cli.Main(["send", "sess-flags"], …) -> error; requires message
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags"}
	return nil
}
```
