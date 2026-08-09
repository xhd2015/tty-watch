# Scenario

**Feature**: empty command after `--` is an error

```
Args [empty-sess --] -> ParseServeArgv -> error (missing command)
```

## Steps

1. Set sid and `--` with no following command tokens.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{"empty-sess", "--"}
	return nil
}
```
