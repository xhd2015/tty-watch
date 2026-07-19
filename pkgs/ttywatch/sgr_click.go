package ttywatch

import "fmt"

// EncodeSGRClick encodes an SGR mouse click sequence.
// row/col are 0-based CLI coordinates; the wire format uses 1-based col/row.
// When release is true, appends press then release sequences.
//
// Wire format:
//
//	press:   ESC [ < btn ; col+1 ; row+1 M
//	release: ESC [ < btn ; col+1 ; row+1 m
func EncodeSGRClick(row, col, btn int, release bool) []byte {
	// CSI SGR mouse: \x1b[<btn;col;rowM (press) / m (release)
	press := []byte(fmt.Sprintf("\x1b[<%d;%d;%dM", btn, col+1, row+1))
	if !release {
		return press
	}
	rel := []byte(fmt.Sprintf("\x1b[<%d;%d;%dm", btn, col+1, row+1))
	out := make([]byte, 0, len(press)+len(rel))
	out = append(out, press...)
	out = append(out, rel...)
	return out
}
