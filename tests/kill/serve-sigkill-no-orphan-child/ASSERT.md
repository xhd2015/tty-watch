## Expected

- Detach starts successfully (`ExitCode == 0`).
- After `SIGKILL` of the `__serve__` process, the PTY **command child** is
  **not** still running (`CommandAlive == false`).
- This frees the OS PTY slave slot even when the child ignored `SIGHUP`.

```go
import (
	"testing"

	"github.com/xhd2015/doctest/session"
)

func Assert(t *testing.T, d *session.Doctest, req *Request, resp *Response, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if resp.ExitCode != 0 {
		t.Fatalf("detach exit %d combined %q", resp.ExitCode, resp.Combined)
	}
	if resp.SessionRunning {
		t.Fatalf("serve still alive after SIGKILL for session %s", resp.SessionID)
	}
	if resp.CommandAlive {
		t.Fatalf("PTY child still alive after serve SIGKILL (orphan leak): session=%s command_pid=%d",
			resp.SessionID, resp.CommandPID)
	}
}
```
