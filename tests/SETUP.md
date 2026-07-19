# Scenario

**Feature**: standalone tty-watch is SSoT for CLI + pkgs/ttywatch — e2e binary phases
and in-process `cli.Main` / `cli.Run` / `cli.ParseArgs` / `EncodeSGRClick` contracts

```
# session-cached binary + isolated home (e2e)
harness -> build ./cmd/tty-watch -> TTY_WATCH_HOME per leaf
tty-watch -> embedded ptywrap + registry -> run/list/watch/attach/snapshot/kill/send
doctest <- Response: exit codes, streams, registry, inject capture

# in-process contract (Mode=cli|run-config|parse-args|encode)
Caller -> cli.Main / cli.Run / cli.ParseArgs / EncodeSGRClick
doctest <- Response: Err/ErrMsg, streams, Bytes, ParsedConfig
```

## Preconditions

- Module path: `github.com/xhd2015/tty-watch`.
- `go` in PATH to build `./cmd/tty-watch` (and lockholder helper).
- Binary entry: `cmd/tty-watch` → `cli.Main(...) error` + `Error:` + `os.Exit(1)`.
- E2E leaves: isolated `TTY_WATCH_HOME` per test; registry at
  `$TTY_WATCH_HOME/registry/*.json`.
- Contract leaves (`cli-main/`, `run-config/`, `send/*/validation`, encode, …)
  set `req.Mode` and do **not** depend on live sessions for flag validation.
- Harness package: `github.com/xhd2015/tty-watch/ttywatchtest`.

## Steps

1. Root Setup builds `tty-watch` once (process-cached) and assigns isolated home.
2. Grouping Setup narrows Phase / Mode / argv / Config.
3. Leaf Setup sets concrete Phase or Mode fields.
4. `Run` dispatches: Mode contract path **or** phase e2e via `ttywatchtest`.
5. Leaf Assert checks Err/ExitCode, streams, registry, inject bytes, or wire bytes.

## Context

- Detach key: Ctrl-] (`\x1d`). Ctrl-C / SIGINT detach observers per product rules.
- Flag validation for send click/query must run **before** registry lookup.
- Ownership: this tree is the SSoT for tty-watch CLI tests (agent-pro trees remain
  until P3; do not rewrite agent-pro product in P1).

```go
import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/xhd2015/tty-watch/ttywatchtest"
)

func Setup(t *testing.T, req *Request) error {
	if _, err := exec.LookPath("go"); err != nil {
		t.Skipf("skipping: go not found in PATH")
	}
	// Always populate Bin/Home so e2e grouping SETUPs that require Bin succeed.
	// Contract Mode leaves ignore them.
	req.Bin = buildTTYWatch(t)
	req.TTYWatchHome = isolatedHome(t)
	return nil
}

func buildTTYWatch(t *testing.T) string {
	return ttywatchtest.BuildTTYWatch(t)
}

func isolatedHome(t *testing.T) string {
	return ttywatchtest.IsolatedHome(t)
}

// assertOK fails when the product API returned a non-nil error (contract success leaves).
func assertOK(t *testing.T, resp *Response) {
	t.Helper()
	if resp == nil {
		t.Fatal("nil response")
	}
	if resp.Err != nil {
		t.Fatalf("expected nil error, got %v\nErrMsg:%s\nstdout:%s\nstderr:%s",
			resp.Err, resp.ErrMsg, resp.Stdout, resp.Stderr)
	}
}

// assertErr fails when the product API returned nil (contract error-path leaves).
func assertErr(t *testing.T, resp *Response) {
	t.Helper()
	if resp == nil {
		t.Fatal("nil response")
	}
	if resp.Err == nil {
		t.Fatalf("expected non-nil error, got nil\nstdout:%s\nstderr:%s\ncombined:%s",
			resp.Stdout, resp.Stderr, resp.Combined)
	}
}

// errorText prefers the returned error message, then stderr, then combined.
func errorText(resp *Response) string {
	if resp == nil {
		return ""
	}
	if strings.TrimSpace(resp.ErrMsg) != "" {
		return resp.ErrMsg
	}
	if resp.Err != nil {
		return resp.Err.Error()
	}
	if strings.TrimSpace(resp.Stderr) != "" {
		return resp.Stderr
	}
	return resp.Combined
}

// assertRegistryHasSession fails when session id is missing from registry dir.
func assertRegistryHasSession(t *testing.T, home, sessionID string) {
	t.Helper()
	if !ttywatchtest.RegistryExists(home, sessionID) {
		t.Fatalf("registry missing %s under %s", sessionID, ttywatchtest.RegistryDir(home))
	}
}

// assertNoHostSessionID fails when combined host output leaks session-N id lines.
func assertNoHostSessionID(t *testing.T, combined string) {
	t.Helper()
	for _, line := range strings.Split(combined, "\n") {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "session-") {
			t.Fatalf("host leaked session id on stdout/stderr: %q", trim)
		}
	}
}

// assertNonZeroExit fails when exit code is 0 for e2e error-path leaves.
func assertNonZeroExit(t *testing.T, resp *Response) {
	t.Helper()
	if resp.ExitCode == 0 {
		t.Fatalf("expected non-zero exit, got 0\nstdout:%s\nstderr:%s\ncombined:%s",
			resp.Stdout, resp.Stderr, resp.Combined)
	}
}

// processArgsLine returns the ps(1) args field for the registry owner PID.
func processArgsLine(t *testing.T, home, sessionID string) (string, error) {
	t.Helper()
	entry, err := ttywatchtest.ReadRegistryEntry(home, sessionID)
	if err != nil {
		return "", err
	}
	if entry.PID <= 0 {
		t.Fatalf("registry %s has invalid pid %d", sessionID, entry.PID)
	}
	out, err := exec.Command("ps", "-p", fmt.Sprintf("%d", entry.PID), "-o", "args=").Output()
	if err != nil {
		return "", fmt.Errorf("ps serve child pid %d: %w", entry.PID, err)
	}
	return strings.TrimSpace(string(out)), nil
}

func phaseRunHeadlessPrintsSessionID(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessPrintsSessionID(t, toPhaseReq(req)))
}

func phaseRunHeadlessNoAttachOutput(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessNoAttachOutput(t, toPhaseReq(req)))
}

func phaseRunHeadlessWaitsUntilChildExits(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessWaitsUntilChildExits(t, toPhaseReq(req)))
}

func phaseRunHeadlessCtrlCForwardsExits(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessCtrlCForwardsExits(t, toPhaseReq(req)))
}

func phaseRunHeadlessCtrlCWaitingLogs(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessCtrlCWaitingLogs(t, toPhaseReq(req)))
}

func phaseRunHeadlessSessionLiveWhileWaiting(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessSessionLiveWhileWaiting(t, toPhaseReq(req)))
}

func phaseRunHeadlessWithCustomSessionID(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessWithCustomSessionID(t, toPhaseReq(req)))
}

func phaseRunHeadlessStopOnFirstArg(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunHeadlessStopOnFirstArg(t, toPhaseReq(req)))
}

func phaseRunDetachPrintsSessionID(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachPrintsSessionID(t, toPhaseReq(req)))
}

func phaseRunDetachNoAttachOutput(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachNoAttachOutput(t, toPhaseReq(req)))
}

func phaseRunDetachSessionSurvivesInList(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachSessionSurvivesInList(t, toPhaseReq(req)))
}

func phaseRunDetachWithCustomSessionID(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachWithCustomSessionID(t, toPhaseReq(req)))
}

func phaseRunDetachMutuallyExclusiveWithHeadless(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachMutuallyExclusiveWithHeadless(t, toPhaseReq(req)))
}

func phaseRunDetachStopOnFirstArg(t *testing.T, req *Request) (*Response, error) {
	return mapPhaseResp(ttywatchtest.PhaseRunDetachStopOnFirstArg(t, toPhaseReq(req)))
}

func toPhaseReq(req *Request) *ttywatchtest.Request {
	if req == nil {
		return &ttywatchtest.Request{}
	}
	return &ttywatchtest.Request{
		Phase:           req.Phase,
		Bin:             req.Bin,
		TTYWatchHome:    req.TTYWatchHome,
		SessionID:       req.SessionID,
		CustomSessionID: req.CustomSessionID,
		RunCommand:      append([]string(nil), req.RunCommand...),
		Detach:          req.Detach,
		SendCtrlC:       req.SendCtrlC,
		Background:      req.Background,
		WatchProbe:      req.WatchProbe,
		SnapshotID:      req.SnapshotID,
		KillID:          req.KillID,
		SendID:          req.SendID,
		SendMessage:     req.SendMessage,
		AttachID:        req.AttachID,
		AttachInput:     req.AttachInput,
		AttachInputB:    req.AttachInputB,
		AttachProbe:     req.AttachProbe,
		Click:           req.Click,
		QueryCursor:     req.QueryCursor,
		ClickRow:        req.ClickRow,
		ClickCol:        req.ClickCol,
		HasClickRow:     req.HasClickRow,
		HasClickCol:     req.HasClickCol,
		Mouse:           req.Mouse,
		HasMouse:        req.HasMouse,
		NoRelease:       req.NoRelease,
		JSON:            req.JSON,
		SendTextArgs:    append([]string(nil), req.SendTextArgs...),
		EncodeRow:       req.EncodeRow,
		EncodeCol:       req.EncodeCol,
		EncodeBtn:       req.EncodeBtn,
		EncodeRelease:   req.EncodeRelease,
	}
}

func mapPhaseResp(r *ttywatchtest.Response, err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	if r == nil {
		return &Response{}, nil
	}
	resp := &Response{
		ExitCode:                           r.ExitCode,
		Stdout:                             r.Stdout,
		Stderr:                             r.Stderr,
		Combined:                           r.Combined,
		SessionID:                          r.SessionID,
		RegistryExists:                     r.RegistryExists,
		RegistryIDs:                        append([]string(nil), r.RegistryIDs...),
		ListOutput:                         r.ListOutput,
		SessionRunning:                     r.SessionRunning,
		SnapshotText:                       r.SnapshotText,
		ContainsEscape:                     r.ContainsEscape,
		TimedOut:                           r.TimedOut,
		Elapsed:                            r.Elapsed,
		GrokModesSeen:                      r.GrokModesSeen,
		TTYCleanupOnDetach:                 r.TTYCleanupOnDetach,
		PostDetachOutput:                   r.PostDetachOutput,
		SourceCheckOK:                      r.SourceCheckOK,
		SourceCheckNote:                    r.SourceCheckNote,
		StdinRestoredBeforeCleanup:         r.StdinRestoredBeforeCleanup,
		KittyPopCleanupInSrc:               r.KittyPopCleanupInSrc,
		AttachStdoutWriterNormalizesRawTTY: r.AttachStdoutWriterNormalizesRawTTY,
		InjectedBytes:                      append([]byte(nil), r.InjectedBytes...),
		Bytes:                              append([]byte(nil), r.InjectedBytes...),
		AltExitCode:                        r.AltExitCode,
		AltStderr:                          r.AltStderr,
		AttachOutput:                       r.AttachOutput,
		AttachBOutput:                      r.AttachBOutput,
		WatchOutput:                        r.WatchOutput,
		RunOutput:                          r.RunOutput,
		LockPath:                           r.LockPath,
		LockHolderPID:                      r.LockHolderPID,
		LockHolderMarker:                   r.LockHolderMarker,
	}
	if r.ExitCode != 0 {
		msg := strings.TrimSpace(r.Stderr)
		if msg == "" {
			msg = strings.TrimSpace(r.Combined)
		}
		if msg == "" {
			msg = fmt.Sprintf("exit %d", r.ExitCode)
		}
		resp.Err = fmt.Errorf("%s", msg)
		resp.ErrMsg = msg
	}
	return resp, nil
}
```
