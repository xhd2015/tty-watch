# Scenario

**Feature**: `send` with zero trailing args is an error

```
cli.Main(["send"], …) -> error; mentions session + message
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send"}
	return nil
}
```
