# Scenario

**Feature**: `tty-watch send --query-cursor` reports host VT cursor (no inject)

```
# host VT cursor only (same model as CPR auto-reply); do NOT inject ESC[6n into child
tty-watch send <id> --query-cursor [--json]
  -> read host screen cursor -> stdout row=R col=C | {"row":R,"col":C}
```

## Preconditions

- **Success leaves** (`plain-at-cup`, `json-at-cup`): phase `send-query-cursor` with detached CUP fixture `printf '\033[5;3H'; sleep 300` → 0-based **row=4 col=2**. Requires working `run --detach` / serve (env SIGKILL of serve is infra).
- **Validation leaves** (`with-text`, `with-click`): phase `send-click-validation`, dummy sid `sess-flags`, **no live session**. Flag parse must run before registry lookup.
- Exclusive with free-text and with `--click`.
- Errors: stderr text, non-zero exit (not JSON bodies in v1).

## Steps

1. Success leaves: `Phase=send-query-cursor`, `QueryCursor=true`, optional `JSON`.
2. Validation leaves: override to `Phase=send-click-validation` + conflict flags.
3. Assert checks stdout geometry or flag-error path.

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	req.Phase = "send-query-cursor"
	req.QueryCursor = true
	return nil
}
```
