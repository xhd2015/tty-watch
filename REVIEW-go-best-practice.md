# Go best-practice review: tty-watch

**Date:** 2026-08-06  
**Scope:** codebase structure, CLI design, flags handling, package layout  
**Lens:** [go-best-practice](https://github.com/xhd2015) topics — `flags-parsing` (+ subcommand/types/cut/collect), `cli` (+ color/streaming/dry-run), `cmd-exec`, `kool-create`, `go-embed-assets`  
**Method:** read `cmd/`, `cli/`, `pkgs/`, README, doctests; built binary and probed live CLI behavior. **No production fixes applied** in this pass.

---

## Executive summary

tty-watch is already a strong Go CLI library hybrid: thin `cmd/tty-watch`, config-driven `cli.ParseArgs` / `cli.Run` / `cli.Main`, intentional use of `less-flags` with `StopOnFirstArg` for `run`, and a clear separation from `pkgs/ttywatch`. Doctests cover many behavioral edges.

The largest gaps vs go-best-practice are:

1. **Per-level `--help` is missing** (core `flags-parsing/subcommand` rule) — several subcommands mis-handle `--help` as a session id or unknown flag.
2. **`run --headless` does not propagate the PTY child’s exit code** (CLI UX / scripting contract), and the thin main always maps errors to exit 1.
3. **Public `cli` package surface leaks registry helpers** via `adapters.go`, blurring the documented “prefer `pkgs/ttywatch`” boundary.

Secondary: prefer `agentdriver.Driver` / `DefaultSelf` over bare `os.Args[0]`; trim dead re-exports; optional polish (`--version`, root help pointer).

**Not applicable (or only lightly):** `kool-create` (existing specialized tool), `go-embed-assets` (no UI/extension assets), `cli/dry-run` (no bulk side-effect pipeline that needs a dry-run gate), `cli/color` (plain tables/text today), `cli/inline-tui-mouse` (host TUI mouse origin; this CLI injects SGR into the *session*, different problem).

---

## What’s already good

| Area | Observation | Topic |
|------|-------------|--------|
| Entry layout | `cmd/tty-watch/main.go` only wires IO + exit; logic lives in importable `cli` | CLI packaging |
| Config-driven API | `ParseArgs` → `Config` → `Run`; IO injected; no `os.Exit` in library | `cli`, testability |
| Subcommand split | Word commands (`run`, `list`, `watch`, …); root has no global flags | `flags-parsing/subcommand` “no toplevel flags” |
| `run` flags | `less-flags` + `StopOnFirstArg()` so command argv is never re-parsed as flags; doctests for flag-like first tokens | `flags-parsing/subcommand`, `StopOnFirstArg` |
| `send` types | `**int` / `*int` style via `rowPtr`/`colPtr` to distinguish unset vs zero | `flags-parsing/types` |
| Mode exclusivity | `--click` vs `--query-cursor` vs text validated in parse path | CLI UX |
| Streaming | `watch`/`attach` stream PTY output; `list` buffers only for a small aligned table (justified) | `cli/streaming` |
| External process control | Detached serve re-exec uses `os/exec` + `Setsid` / wait / kill — correct; **not** a fit for `xgo/support/cmd` Debug printers | `cmd-exec` (when *not* to use) |
| Tests | Large doctest tree; programmatic `cli.Main` / `cli.Run` leaves | practice |
| Docs | README documents layout, env, `StopOnFirstArg`, programmatic API | — |

---

## Findings (severity-ordered)

### Critical

#### C1. Subcommands do not support `-h` / `--help` (and often mis-parse them)

**Topic:** `flags-parsing`, `flags-parsing/subcommand`  
**Rule:** *Every command level needs `--help` with that level’s usage.* Root help alone is not enough. Empty args at a dispatch level → that level’s help.

**Live behavior** (built `./cmd/tty-watch`):

| Invocation | Result |
|------------|--------|
| `tty-watch --help` | Root usage (OK) |
| `tty-watch run --help` / `run -h` | `Error: unrecognized flag: --help` / `-h` (exit 1) |
| `tty-watch send --help` | `send: requires <session-id> and <message>` |
| `tty-watch send demo --help` | `send: unrecognized flag: --help` |
| `tty-watch list --help` | `list: unexpected arguments [--help]` |
| `tty-watch watch\|attach\|kill\|snapshot --help` | Treated as session id → `tty-watch session --help not found` (or registry error) |
| `tty-watch run help` | Tries to exec program `help` (serve re-exec path) |

**Root help** also omits: `Run tty-watch <command> --help for command-specific options.`

**Code:**

- Root: manual `-h`/`--help`/`help` / empty args in `cli/main.go` — OK for “no toplevel flags”.
- `parseRunArgs` (`cli/run.go`): `less-flags` chain has **no** `Help` / `HelpNoExit`.
- `parseSendArgs` (`cli/send.go`): same.
- Positional-only commands: string checks only; no help branch.

**Recommended change:**

1. Per subcommand help strings (usage + flags + examples for `run` / `send`).
2. Flaggy subcommands (`run`, `send`):

   ```go
   remain, err := lessflags.String("--session-id", &sessionIDPtr).
       Bool("--headless", &opts.headless).
       Bool("--detach", &opts.detach).
       Help("-h,--help", runHelp).
       HelpNoExit(). // library must not os.Exit
       StopOnFirstArg().
       Parse(args)
   if errors.Is(err, lessflags.ErrHelp) {
       // Parse already printed help when using Help+HelpNoExit? Check less-flags:
       // Help prints then returns ErrHelp under HelpNoExit — map to nil from ParseArgs
       // after ensuring help text is written to the intended writer if needed.
       return opts, errHelp // or special-case in ParseArgs/Main
   }
   ```

   Prefer **`HelpNoExit()`** because `cli` is a library (`ParseArgs` must not `os.Exit`). Map `ErrHelp` to success + printed help (or a dedicated `Command: "help"` with text). Confirm exact print target in less-flags v1.0.2 (stdout vs stderr).

3. Positional commands (`list`, `watch`, `attach`, `snapshot`, `kill`): if `len(args)==0` or first arg is `-h`/`--help`/`help`, print that command’s usage and return nil (do **not** look up session `"--help"`).
4. Root help: add the “`<command> --help`” pointer; optionally list flags only on leaf help texts.
5. Doctests: one leaf per command for `… --help` exit 0 + usage keywords.

**Severity rationale:** Users and scripts discover CLIs via `--help`. Current errors are wrong and sometimes look like “session missing”.

---

#### C2. Headless child exit status is not a real process exit code

**Topic:** CLI UX (scripting contract); related error-model hygiene  

**Observed:**

```text
tty-watch run --headless -- false          → exit 0
tty-watch run --headless -- bash -c 'exit 42' → exit 0
```

**Why:**

1. `pkgs/ttywatch.ServeSession` waits on the PTY child with `_ = mgr.Wait(...)` and **always returns `nil`** on normal child completion (`server.go`), so the serve process exits 0 regardless of the child’s code.
2. `WaitHeadless` / `exitStatusFromWait` only see the **serve** process status → almost always success.
3. `cli.exitCodeFromWait` builds `cliExitError{code}`, but `cmd/tty-watch/main.go` always does `os.Exit(1)` on *any* non-nil error — never unwraps exit codes. Even if (1) were fixed, main would still flatten codes to 1 and print `Error: exit status N`.

**Recommended change:**

1. **Serve path:** capture PTY exit code from `mgr.Wait` (or equivalent) and exit the serve process with that code (or write it to registry / a small status file if exit-code-through-process is undesirable).
2. **Headless wait:** surface that code as `ttywatch.ExitStatus` / exported exit-code error.
3. **Binary main:**

   ```go
   if err := cli.Main(...); err != nil {
       var code *cli.ExitError // export a typed exit code from cli
       if errors.As(err, &code) {
           os.Exit(code.Code) // no "Error:" prefix for pure exit codes
       }
       fmt.Fprintln(os.Stderr, "Error:", err)
       os.Exit(1)
   }
   ```

4. Document in README: headless parent exit status mirrors the session command.
5. Doctest: `run --headless -- false` → non-zero; `exit 42` → 42 if you choose full propagation.

**Severity rationale:** `run --headless` is the automation-facing mode; silent always-0 breaks shell pipelines and CI.

---

### High

#### H1. `cli` public API re-exports registry / render helpers (`adapters.go`)

**Topic:** package layout / API boundaries (library design adjacent to go-best-practice packaging)

**Issue:** `cli` exports `ReadRegistry`, `WriteRegistry`, `RemoveRegistry*`, `Reserve*SessionID`, `ListRegistryEntries`, `SanitizeForPrint`, `RegistryEntry`, `TTYWatchHome`, etc. Many are **unused outside `adapters.go`** itself (`WriteRegistry`, `SanitizeForPrint`, `Reserve*`, snapshot VT helpers, …). README already says lower-level helpers live in `pkgs/ttywatch`.

Exported dual APIs invite consumers to depend on the wrong package and freeze accidental surface.

**Recommended change:**

1. Keep **exported** `cli` surface to: `Main`, `ParseArgs`, `Run`, `Config`, option structs, `SendMode*`.
2. Unexport or delete unused wrappers; call `ttywatch.*` directly from command files.
3. If some wrappers remain for internal convenience only, make them lowercase (`readRegistry`, …).
4. Optionally add a short `// Package cli` doc comment listing the intentional public surface.

---

#### H2. `run` uses deprecated `BinaryPath: os.Args[0]` instead of `agentdriver.Driver`

**Topic:** external commands / re-exec robustness (`cmd-exec` spirit; project already has `agentdriver`)

**Code:** `cli/run.go` → `HeadlessRunOptions{ BinaryPath: os.Args[0], ... }`.

`HeadlessRunOptions` documents **prefer `Driver` over `BinaryPath`**. `agentdriver.DefaultSelf` / `Resolve` use `os.Executable()`, `LookPath`, and symlink eval — more reliable than a possibly relative `os.Args[0]` (and avoids `go run` temp binary footguns when hosts re-exec).

**Recommended change:**

```go
result, err := ttywatch.HeadlessRun(ctx, ttywatch.HeadlessRunOptions{
    Home:    home,
    SessionID: opts.SessionID,
    Command: opts.Command,
    Driver:  agentdriver.Driver{}, // empty → DefaultSelf inside Resolve
    KeepAlive: false,
})
```

Or pass an explicit resolved driver once. Drop `BinaryPath` at this call site. Add a note for embedding hosts (agent-pro / spl) to set `Driver{Binary, Args}`.

---

#### H3. Root dispatch is manual; fine, but help/error consistency is uneven

**Topic:** `flags-parsing/subcommand` “no toplevel flags”

Manual first-arg switch is **correct** when the root has no flags. Gaps:

- Unknown subcommand: `unknown subcommand %q` — good.
- Internal serve token branch before the switch — good, but must stay first.
- No shared helper for “session-id positional + optional `--help`”.
- `run help` / bare `help` as command name confusion only after C1 is fixed (positional `help` under `run` is a command name, not CLI help — document that).

**Recommended change:** after C1, extract small helpers:

- `parseSessionArg(cmd, args) (id string, err error)` with help handling  
- Keep serve-token check first forever

---

### Medium

#### M1. Custom `duplicateSessionIDFlag` pre-scan

**Topic:** `flags-parsing` / less-flags completeness

`parseRunArgs` walks argv manually for duplicate `--session-id` before `less-flags.Parse`. That works with `StopOnFirstArg` (stops at first non-flag), but is parallel parsing logic that can drift (e.g. if short aliases or `--` rules change).

**Recommended change:**

- Prefer less-flags behavior if it can error on duplicate string flags; or  
- Document why last-wins is rejected; keep one small helper with a unit/doctest; or  
- Accept last-wins (common CLI default) and delete the helper if product agrees.

---

#### M2. No `--version` / build identity

**Topic:** CLI UX (common companion to help)

`tty-watch --version` → `unknown subcommand "--version"`.

**Recommended change:** treat `-v` carefully (conflicts with verbose elsewhere — here unused). Prefer long form only:

```text
tty-watch --version → tty-watch <module version or ldflags>
```

Wire via root manual check (no toplevel less-flags) or a `version` subcommand. Optional `runtime/debug.ReadBuildInfo()`.

---

#### M3. `Config.Home` only programmatic; env documented, no CLI flag

**Topic:** `flags-parsing` (optional), CLI UX

Home override: `TTY_WATCH_HOME` env + `Config.Home` in `Run`. No `--home` flag. Fine for test isolation via env; worth documenting as intentional (avoid flag sprawl — “less flags” spirit).

**Recommended change:** docs-only unless operators need flag form; if added, put on **root** with `StopOnFirstArg` *or* only on mutating commands — prefer env-only to stay lean.

---

#### M4. Diagnostic `exec.Command` for `lsof` / `ps` / `pgrep`

**Topic:** `cmd-exec`

`registry_lock_diag.go` and headless SIGINT helpers use raw `os/exec`. Acceptable: output is parsed, must not print `[cmd] …` debug lines, often short timeouts.

**Recommended change:** keep raw `exec` for process diagnostics. Optionally centralize `exec.CommandContext` + timeout + env in a tiny internal helper; **do not** force `xgo/support/cmd` Debug mode here.

---

#### M5. Large flat `pkgs/ttywatch` package

**Topic:** package layout (maintainability)

~4.5k lines, many concerns: registry, headless re-exec, attach/WS, snapshot/VT render, SGR, observer, reclaim, naming.

Not a go-best-practice recipe violation, but slows review and encourages `cli` adapters.

**Recommended change (incremental, not big-bang):**

- `pkgs/ttywatch/registry`, `…/session`, `…/snapshot`, `…/attach` — only if import cycles allow; or  
- Keep one package, split **files** by domain (already partially done) and export a narrower “public” set in docs.

---

#### M6. `pkgs/` directory prefix

**Topic:** layout convention (`kool-create` / common Go layouts)

Module uses `pkgs/ttywatch` and `pkgs/agentdriver` instead of top-level `ttywatch` / `agentdriver` or `internal/…`. Works; import paths are longer. Internal-only code (if any) could use `internal/` to prevent external import.

**Recommended change:** low priority; if breaking imports is OK someday, `github.com/xhd2015/tty-watch/ttywatch` is more idiomatic. Not worth churn while agent-pro depends on current paths.

---

### Low

#### L1. Color / `NO_COLOR`

**Topic:** `cli/color`  

No ANSI styling today. If status highlights or table accents appear later, implement three-mode resolve (`--color` / `--no-color` / auto + `NO_COLOR`). Skip until needed.

#### L2. Dry-run

**Topic:** `cli/dry-run`  

No multi-step mutating pipeline that needs a side-effect gate. `kill` / `send` are inherently live. No `--dry-run` required.

#### L3. `go:embed` assets

**Topic:** `go-embed-assets`  

No web UI or extension bundle. N/A.

#### L4. `kool-create`

**Topic:** `kool-create`  

Scaffolding recipes (react/go-react/server) don’t map to this PTY manager. Layout is custom and appropriate.

#### L5. Streaming of `list`

**Topic:** `cli/streaming`  

Collect-then-`tabwriter` is the justified buffer case (aligned columns, small N). Optional: stream plain TSV/JSON lines under `--json` later for large registries.

#### L6. Error prefix always `"Error:"`

Thin main always prefixes errors. Fine for human CLI; for pure exit-code errors (C2) suppress prefix. Optional: map “session not found” to exit 2 vs usage errors exit 2 / 64 (BSD sysexits) — only if consumers care.

#### L7. Duplicate debug log implementations

`cli/debug_log.go` and `pkgs/ttywatch/debug_log.go` both honor `TTY_WATCH_DEBUG_LOG_FILE`. Consider one package to avoid drift (low priority).

---

## Topic checklist

| Topic | Status | Notes |
|-------|--------|--------|
| `flags-parsing` | Partial | Used well for `run`/`send`; missing Help/HelpNoExit |
| `flags-parsing/subcommand` | Partial | StopOnFirstArg on `run` ✓; per-level help ✗ |
| `flags-parsing/types` | Good | `*int` unset detection on send |
| `flags-parsing/cut` | N/A / optional | `run` uses StopOnFirstArg remain args (correct); Cut would fit only if you switched to `--exec …` style |
| `flags-parsing/collect` | N/A | No parent→child flag forward filter today |
| `cli` / skill-cli | Good shape | Importable Main/Run; not a multi-skill host |
| `cli/color` | N/A for now | |
| `cli/streaming` | Good | |
| `cli/dry-run` | N/A | |
| `cli/inline-tui-mouse` | N/A | Session inject ≠ host inline TUI origin |
| `cmd-exec` | Intentional raw exec | Don’t retrofit Debug cmd for serve/PTY |
| `kool-create` | N/A | |
| `go-embed-assets` | N/A | |

---

## Package layout (current)

```text
cmd/tty-watch/     thin main (Error + Exit)          ✓ keep
cli/               parse + dispatch + Config API     ✓ keep; shrink exports
pkgs/ttywatch/     core library                      ✓ keep; optional split later
pkgs/agentdriver/  re-exec driver                    ✓ use from cli run path
tests/             doctests                          ✓ keep
ttywatchtest/      harness                           large but OK for e2e
```

**Intended dependency direction:**

```text
cmd → cli → pkgs/ttywatch → (ptywrap, agentdriver)
                ↘ agentdriver
```

Avoid: external code importing `cli` only for `ReadRegistry` (H1).

---

## Recommended fix order

1. **C1** — Per-command help + `HelpNoExit` on `run`/`send`; positional help; root help pointer; doctests.  
2. **C2** — Propagate PTY exit code through serve → headless wait → typed error → main `os.Exit(code)`.  
3. **H2** — `Driver` / `DefaultSelf` instead of `BinaryPath: os.Args[0]`.  
4. **H1** — Unexport / delete `adapters.go` public re-exports.  
5. **M2** — `--version` if operators need it.  
6. Docs: README “every subcommand `--help`”, headless exit-code contract, intentional env-only home.

---

## Suggested doctest leaves (when implementing)

| Leaf | Expect |
|------|--------|
| `cli-main/help-root` | already exists — extend for “command --help” pointer |
| `run/help` | `run --help` exit 0, mentions headless/detach/session-id |
| `send/help` | `send --help` exit 0, mentions click / query-cursor |
| `list/help`, `watch/help`, … | usage, not “session not found” |
| `run/headless/exit-code-false` | exit ≠ 0 |
| `run/headless/exit-code-42` | exit 42 (if full propagation) |

---

## Out of scope for this review

- PTY/VT correctness, attach races, kitty cleanup (covered by existing doctests; not go-best-practice topics).  
- Implementing the fixes above (deferred per request).  
- Changing module path or major package renames without a consumer plan.

---

## Verdict

**Solid foundation** (config-driven CLI, less-flags on the hard `run` path, library-friendly entrypoints). **Not yet aligned** with go-best-practice’s strongest CLI rule: **help at every level**, and headless mode is not yet a trustworthy exit-code citizen for scripts. Address C1 and C2 first; then tidy public API and re-exec driver usage.
)
