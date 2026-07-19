# Scenario

**Feature**: empty args print root usage like help (nil error)

```
cli.Main([], …) -> nil error; stdout usage (same family as -h)
```

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Args = []string{}
	return nil
}
```
