## Expected

- `BuiltArgv` starts with `build-full`.
- Contains `--home /data/home`, `--registry-subdir custom-reg`, bare `--keep-alive`.
- Contains each `--extra-path`, `--env`, `--unset-env` pair in order.
- Contains exactly one `--` separator before pure command `agent run --verbose`.
- Pure command tokens appear only after `--` (command flags like `--verbose` are not serve flags).

## Errors

- None from `Run`.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	_ = req
	if err != nil {
		t.Fatalf("BuildServeArgv: unexpected error: %v", err)
	}
	got := resp.BuiltArgv
	want := []string{
		"build-full",
		"--home", "/data/home",
		"--registry-subdir", "custom-reg",
		"--keep-alive",
		"--extra-path", "/x/bin",
		"--extra-path", "/y/bin",
		"--env", "FOO=bar",
		"--env", "BAZ=1",
		"--unset-env", "NOPE",
		"--",
		"agent", "run", "--verbose",
	}
	if !equalStrings(got, want) {
		t.Fatalf("BuiltArgv =\n  %#v\nwant\n  %#v", got, want)
	}
	// Guard: no TTY_WATCH_* invent in argv itself.
	joined := strings.Join(got, " ")
	for _, k := range reexecKnobKeys {
		if strings.Contains(joined, k) {
			t.Fatalf("BuiltArgv must not embed %s: %v", k, got)
		}
	}
}
```
