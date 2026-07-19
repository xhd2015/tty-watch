# Scenario

**Feature**: negative `--row` is rejected

```
cli.Main(["send", "sess-flags", "--click", "--row", "-1", "--col", "1"], …) -> error
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--click", "--row", "-1", "--col", "1"}
	return nil
}
```
