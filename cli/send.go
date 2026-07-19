package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xhd2015/less-flags"
	"github.com/xhd2015/tty-watch/pkgs/ttywatch"
)

// parseSendArgs validates send argv (after "send") into SendOptions.
// Flag validation runs before registry lookup so dummy sid sess-flags works.
func parseSendArgs(args []string) (SendOptions, error) {
	var opts SendOptions

	if len(args) == 0 {
		return opts, fmt.Errorf("send: requires <session-id> and <message>")
	}

	// Session id is the first positional token (must not be a flag).
	if strings.HasPrefix(args[0], "-") {
		return opts, fmt.Errorf("send: requires <session-id> and <message>")
	}
	opts.Session = args[0]
	rest := args[1:]

	var (
		click       bool
		queryCursor bool
		noRelease   bool
		jsonOut     bool
		rowPtr      *int
		colPtr      *int
		mousePtr    *int
	)
	remain, err := lessflags.Bool("--click", &click).
		Bool("--query-cursor", &queryCursor).
		Bool("--no-release", &noRelease).
		Bool("--json", &jsonOut).
		Int("--row", &rowPtr).
		Int("--col", &colPtr).
		Int("--mouse", &mousePtr).
		Parse(rest)
	if err != nil {
		return opts, fmt.Errorf("send: %w", err)
	}

	hasRow := rowPtr != nil
	hasCol := colPtr != nil
	hasMouse := mousePtr != nil
	hasText := len(remain) > 0

	// Exclusive modes: text | click | query-cursor
	if click && queryCursor {
		return opts, fmt.Errorf("send: --click and --query-cursor are exclusive modes")
	}

	// --row/--col/--mouse/--no-release only valid with --click
	if !click {
		if hasRow || hasCol {
			return opts, fmt.Errorf("send: --row/--col require --click")
		}
		if hasMouse {
			return opts, fmt.Errorf("send: --mouse requires --click")
		}
		if noRelease {
			return opts, fmt.Errorf("send: --no-release requires --click")
		}
	}

	if queryCursor {
		if hasText {
			return opts, fmt.Errorf("send: cannot mix free-text message with --query-cursor")
		}
		opts.Mode = SendModeQueryCursor
		opts.JSON = jsonOut
		return opts, nil
	}

	if click {
		if hasText {
			return opts, fmt.Errorf("send: cannot mix free-text message with --click")
		}
		if !hasRow {
			return opts, fmt.Errorf("send: --click requires --row")
		}
		if !hasCol {
			return opts, fmt.Errorf("send: --click requires --col")
		}
		if *rowPtr < 0 {
			return opts, fmt.Errorf("send: invalid negative --row %d", *rowPtr)
		}
		if *colPtr < 0 {
			return opts, fmt.Errorf("send: invalid negative --col %d", *colPtr)
		}
		opts.Mode = SendModeClick
		opts.Row = *rowPtr
		opts.Col = *colPtr
		if hasMouse {
			opts.Mouse = *mousePtr
		}
		opts.NoRelease = noRelease
		opts.JSON = jsonOut
		return opts, nil
	}

	// Text mode
	if jsonOut {
		return opts, fmt.Errorf("send: --json is only valid with --click or --query-cursor (not text mode)")
	}
	if !hasText {
		return opts, fmt.Errorf("send: requires <session-id> and <message>")
	}
	opts.Mode = SendModeText
	opts.Message = strings.Join(remain, " ")
	return opts, nil
}

func runSend(cfg Config) error {
	if cfg.Send == nil {
		return fmt.Errorf("send: requires <session-id> and <message>")
	}
	opts := cfg.Send

	// Programmatic click validation: session set, row/col ≥ 0.
	if opts.Mode == SendModeClick {
		if strings.TrimSpace(opts.Session) == "" {
			return fmt.Errorf("send: requires <session-id> and <message>")
		}
		if opts.Row < 0 {
			return fmt.Errorf("send: invalid negative row %d", opts.Row)
		}
		if opts.Col < 0 {
			return fmt.Errorf("send: invalid negative col %d", opts.Col)
		}
	}

	home, err := resolveHome(cfg)
	if err != nil {
		return err
	}
	entry, err := ReadRegistry(home, opts.Session)
	if err != nil {
		return err
	}
	if !tcpReachable(entry.ListenAddr) {
		RemoveRegistryIfMatch(home, opts.Session, entry.ListenAddr, entry.PID)
		return fmt.Errorf("tty-watch session %s not found", opts.Session)
	}

	switch opts.Mode {
	case SendModeText:
		if err := prepareSessionInjectMode(entry.ListenAddr, opts.Session); err != nil {
			return err
		}
		return ttywatch.InjectInput(entry.ListenAddr, opts.Session, []byte(opts.Message))

	case SendModeClick:
		release := !opts.NoRelease
		payload := ttywatch.EncodeSGRClick(opts.Row, opts.Col, opts.Mouse, release)
		if err := prepareSessionInjectMode(entry.ListenAddr, opts.Session); err != nil {
			return err
		}
		if err := ttywatch.InjectInput(entry.ListenAddr, opts.Session, payload); err != nil {
			return err
		}
		if opts.JSON {
			ack := struct {
				OK      bool `json:"ok"`
				Row     int  `json:"row"`
				Col     int  `json:"col"`
				Mouse   int  `json:"mouse"`
				Release bool `json:"release"`
			}{
				OK:      true,
				Row:     opts.Row,
				Col:     opts.Col,
				Mouse:   opts.Mouse,
				Release: release,
			}
			enc := json.NewEncoder(cfg.Stdout)
			enc.SetEscapeHTML(false)
			if err := enc.Encode(ack); err != nil {
				return err
			}
		}
		return nil

	case SendModeQueryCursor:
		row, col, err := ttywatch.QueryCursor(entry.ListenAddr, opts.Session)
		if err != nil {
			return err
		}
		if opts.JSON {
			out := struct {
				Row int `json:"row"`
				Col int `json:"col"`
			}{Row: row, Col: col}
			enc := json.NewEncoder(cfg.Stdout)
			enc.SetEscapeHTML(false)
			return enc.Encode(out)
		}
		fmt.Fprintf(cfg.Stdout, "row=%d col=%d\n", row, col)
		return nil

	default:
		return fmt.Errorf("send: unknown mode")
	}
}
