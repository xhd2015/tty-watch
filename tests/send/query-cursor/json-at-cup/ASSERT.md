## Expected

- Exit code 0.
- Stdout JSON `{"row":4,"col":2}` (0-based after CUP 5;3).

```go
import (
	"encoding/json"
	"testing"
)

func Assert(t *testing.T, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("query-cursor --json exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	var cur struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}
	if err := json.Unmarshal([]byte(resp.Stdout), &cur); err != nil {
		t.Fatalf("stdout JSON: %v; stdout=%q", err, resp.Stdout)
	}
	if cur.Row != 4 || cur.Col != 2 {
		t.Fatalf("cursor %+v, want row=4 col=2", cur)
	}
}
```
