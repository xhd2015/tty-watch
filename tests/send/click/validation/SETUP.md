# Scenario

**Feature**: click mode rejects incomplete or illegal flag combinations

```
# missing coords / free-text mix / negative / flags without --click
cli.Main(["send", "sess-flags", …], …) -> non-nil error; flag error
# validation BEFORE registry lookup
```

## Preconditions

- Dummy sid always `sess-flags`.
- Implementer contract: parse/validate mode flags before registry so these
  leaves do not depend on `run --detach`.

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Mode = "cli"
	// Leaves override Args; default dummy sid for validation-before-registry.
	if len(req.Args) == 0 {
		req.Args = []string{"send", "sess-flags"}
	}
	return nil
}
```
