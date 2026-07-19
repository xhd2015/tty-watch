# Scenario

**Feature**: free-text args cannot mix with `--click`

```
cli.Main(["send", "sess-flags", "--click", "--row", "1", "--col", "1", "hello"], …)
  -> error; cannot mix free-text with click
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--click", "--row", "1", "--col", "1", "hello"}
	return nil
}
```
