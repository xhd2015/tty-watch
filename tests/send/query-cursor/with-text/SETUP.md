# Scenario

**Feature**: free-text cannot mix with `--query-cursor`

```
cli.Main(["send", "sess-flags", "--query-cursor", "hello"], …)
  -> error; cannot mix free-text with query-cursor
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Mode = "cli"
	req.Args = []string{"send", "sess-flags", "--query-cursor", "hello"}
	return nil
}
```
