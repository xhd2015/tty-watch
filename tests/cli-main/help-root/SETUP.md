# Scenario

**Feature**: `Main([]string{"-h"})` prints root usage and returns nil error

```
cli.Main(["-h"], …) -> nil error; stdout Usage + subcommands (run, send, …)
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{"-h"}
	return nil
}
```
