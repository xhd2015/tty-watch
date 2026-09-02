package ttywatch

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hinshun/vt10x"
)

// grokChangelogGoodScrollback is the final grok changelog screen drawn via absolute CUP.
func grokChangelogGoodScrollback() string {
	var b strings.Builder
	b.WriteString("\x1b[?1049h\x1b[?25l")
	b.WriteString("\x1b[1;1H\x1b[K  Quit ctrl+q    Changelog    Settings")
	b.WriteString("\x1b[3;1H╭──────────────────────────────────────────────────────────╮")
	b.WriteString("\x1b[4;1H│ Grok Build Changelog                                     │")
	b.WriteString("\x1b[5;1H│                                                          │")
	b.WriteString("\x1b[6;1H│ • Snapshot screen-frame parity fix                       │")
	b.WriteString("\x1b[7;1H╰──────────────────────────────────────────────────────────╯")
	b.WriteString("\x1b[20;1H❯ Ask anything")
	b.WriteString("\x1b[24;1HLogged in with API key · Grok Build")
	return b.String()
}

// grokChangelogGarbledScrollback mimics scrollback polluted by changelog-only boot redraws
// that double vt10x replay collapses to Quit q without ctrl+q and no bordered box.
func grokChangelogGarbledScrollback() string {
	var b strings.Builder
	b.WriteString("\x1b[?1049h\x1b[?25l")
	b.WriteString("\x1b[1;1HChangelog")
	b.WriteString("\x1b[2;1HQuit q")
	b.WriteString("\x1b[3;1HGrok Build")
	b.WriteString("\x1b[4;1H• Snapshot screen-frame parity fix")
	b.WriteString("\x1b[20;1HAsk anything")
	return b.String()
}

func grokChangelogWants() []string {
	return []string{"ctrl+q", "Grok Build Changelog", "❯", "Logged in with API key"}
}

func TestRenderSnapshotOutput_prefersFrameOverGarbledScrollback(t *testing.T) {
	good := grokChangelogGoodScrollback()
	frame, ok := renderScreenSnapshotFrame([]byte(good), 80, 24)
	if !ok {
		t.Fatal("failed to build grok changelog screen snapshot frame")
	}
	badScroll := grokChangelogGarbledScrollback()

	scrollOnly := renderSnapshotOutput("", badScroll, 80, 24, false)
	if strings.Contains(scrollOnly, "ctrl+q") {
		t.Fatalf("garbled scrollback fixture should lack ctrl+q, got %q", scrollOnly)
	}

	out := renderSnapshotOutput(string(frame), badScroll, 80, 24, false)
	for _, want := range grokChangelogWants() {
		if !strings.Contains(out, want) {
			t.Fatalf("RenderSnapshotOutput must prefer screen frame over scrollback; missing %q, got %q", want, out)
		}
	}
	if strings.Contains(out, "Quit q") && !strings.Contains(out, "ctrl+q") {
		t.Fatalf("RenderSnapshotOutput used garbled scrollback (Quit q without ctrl+q), got %q", out)
	}
}

// grokPostTurnGoodScrollback mimics a completed grok turn: prompt, assistant reply,
// stop indicator, turn-complete status, and input box — content between UI chrome
// must not be dropped by snapshot rendering.
func grokPostTurnGoodScrollback() string {
	var b strings.Builder
	b.WriteString("\x1b[?1049h\x1b[?25l")
	b.WriteString("\x1b[2;1H  worktree header")
	b.WriteString("\x1b[5;1H     ❯ one word of France captial                                   11:39 AM")
	b.WriteString("\x1b[8;1H     making changes. Looking at the task: one word of France captial.")
	b.WriteString("\x1b[9;1H     Your request has been processed. Everything looks good.")
	b.WriteString("\x1b[14;1H     ◆ stop  [hooks: 1]")
	b.WriteString("\x1b[16;1H     Turn completed in 1.4s.")
	b.WriteString("\x1b[19;1H  ╭──────────────────────────────────────────────────────────╮")
	b.WriteString("\x1b[20;1H  │ ❯                                                        │")
	b.WriteString("\x1b[21;1H  ╰──────────────────────────────────────────────────────────╯")
	b.WriteString("\x1b[23;1H  Shift+Tab:mode  │  Ctrl+.:shortcuts")
	return b.String()
}

func grokPostTurnWants() []string {
	return []string{
		"one word of France captial",
		"processed. Everything looks good",
		"stop  [hooks: 1]",
		"Turn completed in 1.4s.",
		"╭",
		"Ctrl+.:shortcuts",
	}
}

func TestRenderSnapshotOutput_grokPostTurnFrameKeepsConversation(t *testing.T) {
	good := grokPostTurnGoodScrollback()
	frame, ok := renderScreenSnapshotFrame([]byte(good), 80, 24)
	if !ok {
		t.Fatal("failed to build grok post-turn screen snapshot frame")
	}

	out := renderSnapshotOutput(string(frame), "", 80, 24, false)
	for _, want := range grokPostTurnWants() {
		if !strings.Contains(out, want) {
			t.Fatalf("snapshot must keep grok post-turn conversation (watch parity); missing %q, got %q", want, out)
		}
	}
}

func TestScrollbackToScreenText_singlePassMatchesGoodFixture(t *testing.T) {
	good := grokChangelogGoodScrollback()
	direct, ok := screenSnapshotToText([]byte(good), 80, 24)
	if !ok {
		t.Fatal("single-pass screenSnapshotToText failed on good grok fixture")
	}
	for _, want := range grokChangelogWants() {
		if !strings.Contains(string(direct), want) {
			t.Fatalf("single-pass replay missing %q, got %q", want, direct)
		}
	}

	// Double vt10x scrollback replay must not drop bordered box from good fixture.
	out := renderSnapshotOutput("", good, 80, 24, false)
	for _, want := range grokChangelogWants() {
		if !strings.Contains(out, want) {
			t.Fatalf("scrollback render missing %q after replay, got %q", want, out)
		}
	}
}

func TestRenderSnapshotOutput_preservesMidSpanBlankLines(t *testing.T) {
	// CRLF matches PTY ONLCR output from printf "...\n" under a real slave TTY.
	scrollback := "\x1b[2J\x1b[Hline1\r\n\r\nline2\r\n"
	out := renderSnapshotOutput("", scrollback, 80, 24, false)
	if !strings.Contains(out, "line1\n\nline2") {
		t.Fatalf("snapshot must keep mid-span blank between line1 and line2, got %q", out)
	}
}

func TestRenderSnapshotOutput_statusLikeSparseDoesNotGlueShortRows(t *testing.T) {
	// Use CRLF like a real PTY (ONLCR) so rows start at column 0.
	var b strings.Builder
	b.WriteString("\x1b[2J\x1b[H")
	b.WriteString(" codelens  test/id          DEGRADED         16:00:00  refresh 2s\r\n")
	b.WriteString("\r\n")
	b.WriteString(" server   tcp 2/2\r\n")
	b.WriteString(" worker   etcd 3\r\n")
	b.WriteString(" database ok\r\n")
	b.WriteString("\r\n")
	b.WriteString(" pause    sticky on\r\n")
	b.WriteString(" quiet    00:00-04:00\r\n")
	b.WriteString(" optimize enabled\r\n")
	b.WriteString("\r\n")
	b.WriteString(" activity idle\r\n")
	b.WriteString("\r\n")
	b.WriteString(" q quit   r refresh\r\n")

	out := renderSnapshotOutput("", b.String(), 80, 24, false)
	for _, bad := range []string{"2sserver", "okpause", "onquiet", "enabledactivity"} {
		if strings.Contains(out, bad) {
			t.Fatalf("short status rows must not glue (%q present), got %q", bad, out)
		}
	}
	for _, want := range []string{"server   tcp", "pause    sticky on", "quiet    00:00-04:00", "q quit"} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing %q in status-like snapshot: %q", want, out)
		}
	}
	if !(strings.Contains(out, "refresh 2s\n\n") && strings.Contains(out, "server   tcp")) {
		t.Fatalf("expected blank after header before server block, got %q", out)
	}
}

func TestMergeSnapshotWrappedLinesCols_fullWidthMerges(t *testing.T) {
	prev := strings.Repeat("x", 79)
	next := "yyyy"
	out := mergeSnapshotWrappedLinesCols([]string{prev, next}, 80, false)
	if len(out) != 1 || out[0] != prev+next {
		t.Fatalf("full-width wrap must merge, got %#v", out)
	}
}

func TestMergeSnapshotWrappedLinesCols_longSoftWrapWhenColsInflated(t *testing.T) {
	// Content wrapped at 80 but snapshot cols reported as 100.
	prev := "WIDE_LINE_MARKER_" + strings.Repeat("x", 63) // 80 runes
	next := strings.Repeat("x", 12)
	out := mergeSnapshotWrappedLinesCols([]string{prev, next}, 100, false)
	if len(out) != 1 || out[0] != prev+next {
		t.Fatalf("long soft-wrap must merge when cols inflated, got %#v", out)
	}
}

func TestMergeSnapshotWrappedLinesCols_shortRowsDoNotMerge(t *testing.T) {
	out := mergeSnapshotWrappedLinesCols([]string{" pause    sticky on", " quiet    00:00-04:00"}, 80, false)
	if len(out) != 2 {
		t.Fatalf("short consecutive rows must stay separate, got %#v", out)
	}
}

func TestRenderSnapshotTextLine_preserveTrailingUsesCursor(t *testing.T) {
	vt := vt10x.New(vt10x.WithSize(40, 4))
	if _, err := fmt.Fprintf(vt, "› a "); err != nil {
		t.Fatal(err)
	}
	trimmed := renderSnapshotTextLine(vt, 40, 0, false)
	if trimmed != "› a" {
		t.Fatalf("default truncate got %q want %q", trimmed, "› a")
	}
	preserved := renderSnapshotTextLine(vt, 40, 0, true)
	if preserved != "› a " {
		t.Fatalf("preserve got %q want %q (cursor=%+v)", preserved, "› a ", vt.Cursor())
	}
}

func TestScreenSnapshotFrameToText_preserveTrailingSpace(t *testing.T) {
	// Minimal ptywrap-like CUP frame with a trailing space on the composer line.
	var b strings.Builder
	b.WriteString("\x1b[?25l")
	b.WriteString("\x1b[2J")
	b.WriteString("\x1b[1;1H› a ")
	b.WriteString("\x1b[2;1H  gpt-5.6-luna xhigh · ~/ws")
	frame := []byte(b.String())

	def, ok := screenSnapshotFrameToTextPreserve(frame, 80, 24, false)
	if !ok {
		t.Fatal("default CUP render failed")
	}
	if strings.Contains(string(def), "› a ") {
		t.Fatalf("default must trim trailing space, got %q", def)
	}
	if !strings.Contains(string(def), "› a") {
		t.Fatalf("default missing draft, got %q", def)
	}

	pres, ok := screenSnapshotFrameToTextPreserve(frame, 80, 24, true)
	if !ok {
		t.Fatal("preserve CUP render failed")
	}
	if !strings.Contains(string(pres), "› a ") {
		t.Fatalf("preserve must keep trailing space, got %q", pres)
	}
}

// TestRenderSnapshotOutput_usableFrameIgnoresSlidingScrollbackPrefix locks the
// idle SoftExit fix: identical live frames with different scrollback plain
// prefixes (ring head slide after SPACE+DEL) must yield identical SnapshotText.
func TestRenderSnapshotOutput_usableFrameIgnoresSlidingScrollbackPrefix(t *testing.T) {
	var b strings.Builder
	b.WriteString("\x1b[?25l")
	b.WriteString("\x1b[2J")
	b.WriteString("\x1b[1;1H› Ask Codex to do anything")
	b.WriteString("\x1b[2;1H  gpt-5.6-luna max")
	frame := b.String()
	if !isScreenSnapshotFrame([]byte(frame)) {
		t.Fatal("fixture must be a usable screen-snapshot frame")
	}

	sbRest := "old-head-history\nmore\n\x1b[?2026h\x1b[1;1Hignored"
	sbAfterDEL := "new-head-after-ring-slide\n\x1b[?2026h\x1b[1;1Hignored"

	rest := RenderSnapshotOutput(frame, sbRest, 80, 24)
	after := RenderSnapshotOutput(frame, sbAfterDEL, 80, 24)
	if rest != after {
		t.Fatalf("usable frame must ignore sliding scrollback prefix; rest=%q after=%q", rest, after)
	}
	if strings.Contains(rest, "old-head-history") || strings.Contains(rest, "new-head-after-ring-slide") {
		t.Fatalf("scrollback plain prefix must not appear in frame render, got %q", rest)
	}
	if !strings.Contains(rest, "Ask Codex to do anything") {
		t.Fatalf("frame content missing, got %q", rest)
	}

	frameOnly := RenderSnapshotOutput(frame, "", 80, 24)
	if rest != frameOnly {
		t.Fatalf("frame+scrollback must equal frame-only; got %q vs %q", rest, frameOnly)
	}
}

func TestRenderSnapshotOutputOpts_preserveTrailingSpace(t *testing.T) {
	var b strings.Builder
	b.WriteString("\x1b[?25l")
	b.WriteString("\x1b[2J")
	b.WriteString("\x1b[1;1H› a ")
	frame := b.String()

	def := RenderSnapshotOutputOpts(frame, "", 80, 24, SnapshotTextOptions{})
	pres := RenderSnapshotOutputOpts(frame, "", 80, 24, SnapshotTextOptions{PreserveTrailingSpace: true})
	if strings.Contains(def, "› a ") {
		t.Fatalf("default RenderSnapshotOutputOpts trimmed poorly? got %q", def)
	}
	// Live frames prefer VT path; cursor after "› a " must preserve.
	if !strings.Contains(pres, "› a ") && !strings.HasSuffix(strings.TrimRight(pres, "\n"), "› a ") {
		// VT path: writing CUP into vt10x then reading cells — ensure preserve works
		// when cursor lands after the trailing space.
		vt := vt10x.New(vt10x.WithSize(80, 24))
		_, _ = vt.Write([]byte("\x1b[H\x1b[2J› a "))
		line := renderSnapshotTextLine(vt, 80, 0, true)
		if line != "› a " {
			t.Fatalf("preserve path: RenderSnapshotOutputOpts=%q vt line=%q want trailing space", pres, line)
		}
	}
}