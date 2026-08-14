# Scenario

**Feature**: programmatic `cli.Run(Config)` and pure `cli.ParseArgs` (no argv in Run)

```
# Mode=run-config — config-driven execution
Caller -> cli.Run(Config{Command, Send, Stdout, …}) -> error

# Mode=parse-args — flag map only (no registry / no Run)
Caller -> cli.ParseArgs(["send", sid, "--click", …]) -> (Config, error)
```

## Preconditions

- No package-level stream globals; harness attaches buffer IO on Config for Run.
- Run never parses CLI flags — leaves build Config in Setup.
- ParseArgs leaf does not need a live session or TTY_WATCH_HOME.
- Send click Run validation: Mode==Click requires non-empty Session and
  row/col ≥ 0 (use col &lt; 0 for missing-col leaf).

## Steps

1. Leaf sets Mode to `run-config` or `parse-args`.
2. Leaf fills `req.Config` or `req.Args`.
3. Harness Run attaches streams (run-config) or captures ParsedConfig (parse-args).
4. Assert Err/ErrMsg, stdout, or ParsedConfig.Send fields.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	// Default to run-config; parse-args leaf overrides Mode.
	if req.Mode == "" || req.Mode == "cli" {
		req.Mode = "run-config"
	}
	return nil
}
```
