# Scenario

**Feature**: `tty-watch send --click` encodes SGR mouse events and injects them into a live session

```
# click mode
tty-watch send <id> --click --row R --col C [--mouse B] [--no-release] [--json]
  -> validate flags -> EncodeSGRClick -> InjectInput
  -> silent stdout | JSON ack
```

## Preconditions

- Modes are exclusive: click vs free-text vs query-cursor.
- Both `--row` and `--col` required with `--click`.
- `--row` / `--col` / `--mouse` / `--no-release` only valid with `--click`.
- Release default ON; `--no-release` → press only.
- Default `--mouse` = 0.
- No PTY size bounds check.

## Steps

1. Inject leaves set `Phase=send-click-capture` and click fields.
2. Validation leaves set `Phase=send-click-validation` with incomplete/illegal flags.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	_ = d
	req.Click = true
	return nil
}
```
