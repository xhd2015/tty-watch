# Scenario

**Feature**: `--click` and `--query-cursor` cannot combine

```
cli.Main(["send", "sess-flags", "--click", "--row", "1", "--col", "1", "--query-cursor"], …)
  -> error; exclusive modes
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Mode = "cli"
	req.Args = []string{
		"send", "sess-flags",
		"--click", "--row", "1", "--col", "1",
		"--query-cursor",
	}
	return nil
}
```
