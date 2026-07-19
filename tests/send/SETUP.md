# Scenario

**Feature**: `tty-watch send` injects follow-up bytes, SGR mouse clicks, or reports host VT cursor

```
# text mode (unchanged)
tty-watch send <id> <message...> -> ReadRegistry -> POST input -> exit 0 (silent)

# click mode — encode SGR press/release, inject, optional JSON ack
tty-watch send <id> --click --row R --col C [--mouse B] [--no-release] [--json]
  -> EncodeSGRClick(R,C,B,release) -> InjectInput(SGR bytes)
  -> stdout empty | {"ok":true,"row":R,"col":C,"mouse":B,"release":bool}

# query-cursor — host VT cursor only (no child inject of ESC[6n)
tty-watch send <id> --query-cursor [--json]
  -> read host screen cursor -> stdout "row=R col=C" | {"row":R,"col":C}
```

## Preconditions

- Success text/click inject leaves start a detached byte-capture child (`cat > capture.bin`) before send. **Requires working `run --detach` / serve** (env may SIGKILL serve; that is infra, not product assert).
- **Flag-validation leaves** (`send/click/validation/*`, `send/query-cursor/with-text`, `with-click`, `missing-args/*`) use **Mode=cli** + dummy sid `sess-flags` — **no live session**; `cli.Main` validates flags before registry.
- Success `query-cursor` CUP leaves need a live detached session with CUP fixture (same detach requirement as inject).
- Message bytes (text mode) are joined from CLI args after session id with a single space.
- Click/query modes are exclusive of free-text message args and of each other.
- **0-based** `--row` / `--col` on CLI; wire SGR uses 1-based `col+1` / `row+1`.
- Pure wire-format leaves live under `unit/encode-sgr-click/` (Mode=encode → product `ttywatch.EncodeSGRClick`).

## Steps

1. Contract validation leaves set `req.Mode="cli"` + `req.Args`.
2. E2E inject/query leaves set `req.Phase` and mode fields (`Click`, `QueryCursor`, row/col, …).
3. Assert checks `resp.Err` (contract) or exit code / `InjectedBytes` (e2e).

## Context

### Phase keys (send subtree)

| Phase | Purpose |
|-------|---------|
| `send-injects-verbatim` / `send-no-suffix` / `send-preserves-whitespace` | text inject + cat capture |
| `send-missing-args` / `send-missing` / `send-stale` | text error paths |
| `send-click-capture` | `--click` inject + cat capture (`InjectedBytes`); needs detach |
| `send-click-validation` | flag/mode validation; dummy sid `sess-flags`; **no session** |
| `send-query-cursor` | `--query-cursor` success path (CUP 5;3 fixture); needs detach |

### Request fields (click / query)

| Field | CLI |
|-------|-----|
| `Click` | `--click` |
| `QueryCursor` | `--query-cursor` |
| `HasClickRow` + `ClickRow` | `--row <n>` (0-based) |
| `HasClickCol` + `ClickCol` | `--col <n>` (0-based) |
| `HasMouse` + `Mouse` | `--mouse <btn>` (default 0 when omitted) |
| `NoRelease` | `--no-release` |
| `JSON` | `--json` |
| `SendMessage` / `SendTextArgs` | free-text args (text mode or mix-error) |

### Wire format (sealed)

`EncodeSGRClick(row, col, btn, release)` →  
`ESC [ < btn ; col+1 ; row+1 M` then if release `ESC [ < btn ; col+1 ; row+1 m`.

```go
import "testing"

func Setup(t *testing.T, req *Request) error {
	// Contract Mode=cli leaves validate flags without a live session.
	// E2E inject/query leaves need Bin from root Setup.
	if req.Mode == "" && req.Bin == "" {
		t.Fatalf("send setup: tty-watch binary not built (e2e Phase path)")
	}
	return nil
}
```
