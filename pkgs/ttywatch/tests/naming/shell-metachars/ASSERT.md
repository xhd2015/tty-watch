## Expected

- `resp.Slug` is `sh_c_echo_whoami_cat_done_HOME` (metachars `;`, `|`, `&`, `$` removed).
- Slug contains none of `;|&$`\"'<>(){}[]*?~!#` characters.
- `resp.ServeSubcommand` is `__serve_sh_c_echo_whoami_cat_done_HOME__`.

## Errors

- None from `Run`.

```go
import (
	"strings"
	"testing"

	"github.com/xhd2015/doctest/session"
)

const forbiddenMetachars = `;|&$"'<>(){}[]*?~!#`

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	_ = d
	if err != nil {
		t.Fatal(err)
	}
	wantSlug := "sh_c_echo_whoami_cat_done_HOME"
	if resp.Slug != wantSlug {
		t.Fatalf("Slug = %q, want %q", resp.Slug, wantSlug)
	}
	if strings.ContainsAny(resp.Slug, forbiddenMetachars) {
		t.Fatalf("slug still contains shell metacharacters: %q", resp.Slug)
	}
	assertServeWrap(t, resp.Slug, resp.ServeSubcommand)
	assertNoBareServe(t, resp.ServeSubcommand)
}
```