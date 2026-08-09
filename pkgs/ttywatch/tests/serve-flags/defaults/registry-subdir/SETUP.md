# Scenario

**Feature**: DefaultRegistrySubdir is the constant `"registry"`

```
DefaultRegistrySubdir() -> "registry"
# no TTY_WATCH_REGISTRY_SUBDIR getenv
```

## Steps

1. Mode already `defaults`; UserHome unused for this leaf.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Explicit empty UserHome: subdir default is independent of home.
	req.UserHome = ""
	return nil
}
```
