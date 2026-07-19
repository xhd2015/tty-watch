package ttywatch

import (
	"fmt"

	"github.com/hinshun/vt10x"
)

// QueryCursor reads the host VT cursor position for a live session (0-based
// row/col for CLI). Uses the same host screen model as snapshot export — does
// not inject CSI 6n into the child PTY.
func QueryCursor(listenAddr, sessionID string) (row, col int, err error) {
	frame, scrollback, cols, rows, err := ReadSnapshot(listenAddr, sessionID)
	if err != nil {
		return 0, 0, fmt.Errorf("query cursor: %w", err)
	}
	if cols <= 0 {
		cols = 80
	}
	if rows <= 0 {
		rows = 24
	}

	data := []byte(frame)
	if len(data) == 0 {
		data = []byte(scrollback)
	}
	if len(data) == 0 {
		return 0, 0, fmt.Errorf("query cursor: cursor unavailable (empty screen)")
	}

	vt := vt10x.New(vt10x.WithSize(cols, rows))
	if _, werr := vt.Write(data); werr != nil {
		return 0, 0, fmt.Errorf("query cursor: %w", werr)
	}
	cursor := vt.Cursor()
	// vt10x Cursor is 0-based (X=col, Y=row).
	return cursor.Y, cursor.X, nil
}
