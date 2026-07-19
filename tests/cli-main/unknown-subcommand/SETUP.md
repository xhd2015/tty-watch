# Scenario

**Feature**: unknown subcommand returns non-nil error

```
cli.Main(["nope"], …) -> error; message mentions unknown / subcommand
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"nope"}
	return nil
}
```
