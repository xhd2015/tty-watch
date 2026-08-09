# Scenario

**Feature**: all serve flags together parse into ServeArgv

```
Args [sid --home H --registry-subdir S --keep-alive
      --extra-path P1 --extra-path P2
      --env K=V --unset-env U -- agent --flag]
  -> full ServeArgv fields; Command pure agent argv
```

## Steps

1. Set every flag kind once or twice, then `--` and a multi-word command.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Args = []string{
		"full-sess",
		"--home", "/tmp/tty-home-x",
		"--registry-subdir", "grok-tty-registry",
		"--keep-alive",
		"--extra-path", "/opt/bin",
		"--extra-path", "/opt/extra",
		"--env", "AGENT_COLOR=1",
		"--unset-env", "HTTP_PROXY",
		"--",
		"codex", "run", "--model", "gpt-5",
	}
	return nil
}
```
