# Scenario

**Feature**: `--row` / `--col` without `--click` is an error

```
cli.Main(["send", "sess-flags", "--row", "1", "--col", "1"], …)
  -> error; --row/--col require --click
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--row", "1", "--col", "1"}
	return nil
}
```
