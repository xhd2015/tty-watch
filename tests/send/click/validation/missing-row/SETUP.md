# Scenario

**Feature**: `--click --col` without `--row` is an error

```
cli.Main(["send", "sess-flags", "--click", "--col", "1"], …) -> error; mentions row
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--click", "--col", "1"}
	return nil
}
```
