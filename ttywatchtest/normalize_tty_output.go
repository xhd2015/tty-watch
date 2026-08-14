package ttywatchtest

import "bytes"

// NormalizeTTYOutput mirrors pkgs/ttywatch/attach_client.go normalizeTTYOutput for doctest
// regression locks on raw TTY CRLF shaping (bare LF → CRLF; standalone CR preserved;
// blank line before [Terminal exited] collapsed).
func NormalizeTTYOutput(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	var out bytes.Buffer
	out.Grow(len(b) + bytes.Count(b, []byte{'\n'}))
	for i, c := range b {
		if c != '\n' {
			out.WriteByte(c)
			continue
		}
		if i > 0 && b[i-1] == '\r' {
			out.WriteByte('\n')
			continue
		}
		out.WriteString("\r\n")
	}
	return collapseBlankBeforeExitMarker(out.Bytes())
}

func collapseBlankBeforeExitMarker(b []byte) []byte {
	const marker = "[Terminal exited]"
	for {
		n := bytes.ReplaceAll(b, []byte("\r\n\r\n"+marker), []byte("\r\n"+marker))
		n = bytes.ReplaceAll(n, []byte("\n\n"+marker), []byte("\n"+marker))
		if bytes.Equal(n, b) {
			return n
		}
		b = n
	}
}