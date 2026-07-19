# Scenario

**Feature**: `--json` is rejected with text-only send mode

```
cli.Main(["send", "sess-flags", "--json", "hello"], …)
  -> error; --json only for click/query-cursor
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"send", "sess-flags", "--json", "hello"}
	return nil
}
```
