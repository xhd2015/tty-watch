# Scenario

**Feature**: `--click --row` without `--col` is an error

```
cli.Main(["send", "sess-flags", "--click", "--row", "1"], …) -> error; mentions col
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--click", "--row", "1"}
	return nil
}
```
