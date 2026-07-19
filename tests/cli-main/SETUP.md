# Scenario

**Feature**: programmatic `cli.Main` root dispatch (help + unknown)

```
# argv without program name; returns error only (no exit code)
cli.Main(["-h"] | [] | ["nope"], stdin, stdout, stderr) -> error + streams
# help/usage on stdout nil error; unknown -> non-nil error (msg keywords)
```

## Preconditions

- Mode is `cli` (default).
- No registry home or session required.

## Steps

1. Leaf sets `req.Args`.
2. Run invokes `cli.Main(req.Args, …)`.
3. Assert `resp.Err` and stdout/stderr content.

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Mode = "cli"
	return nil
}
```
