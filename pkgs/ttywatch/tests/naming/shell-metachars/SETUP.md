# Scenario

**Feature**: shell metacharacters in argv become underscore separators

```
argv [sh -c echo;whoami|cat&done]
  -> SlugifyCommandLine -> metachars ; | & $ sanitized
  -> ServeSubcommand -> __serve_{slug}__
```

## Steps

1. Set `req.Argv` with a `sh -c` script fragment containing common metacharacters.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Argv = []string{"sh", "-c", "echo;whoami|cat&done$HOME"}
	return nil
}
```