package ttywatch

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRegistryEntry_commandPIDRoundTrip(t *testing.T) {
	dir := t.TempDir()
	cfg := RegistryConfig{Home: dir, Subdir: "codex-tty-registry"}
	entry := RegistryEntry{
		SessionID:       "sess-cmd-1",
		ListenAddr:      "127.0.0.1:9",
		PID:             111,
		CreatedAt:       "2026-08-07T00:00:00Z",
		Command:         []string{"codex"},
		Cwd:             "/tmp",
		CommandPID:      222,
		CommandExited:   true,
		CommandExitedAt: "2026-08-07T00:01:00Z",
	}
	if err := WriteRegistry(cfg, entry); err != nil {
		t.Fatalf("WriteRegistry: %v", err)
	}
	got, err := ReadRegistry(cfg, entry.SessionID)
	if err != nil {
		t.Fatalf("ReadRegistry: %v", err)
	}
	if got.CommandPID != 222 {
		t.Fatalf("CommandPID: got %d want 222", got.CommandPID)
	}
	if !got.CommandExited {
		t.Fatal("CommandExited want true")
	}
	if got.CommandExitedAt != entry.CommandExitedAt {
		t.Fatalf("CommandExitedAt: got %q want %q", got.CommandExitedAt, entry.CommandExitedAt)
	}
	// Raw JSON keys present for operators.
	raw, err := os.ReadFile(filepath.Join(RegistryDir(cfg), entry.SessionID+".json"))
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	if _, ok := m["command_pid"]; !ok {
		t.Fatalf("json missing command_pid: %s", raw)
	}
	if _, ok := m["command_exited"]; !ok {
		t.Fatalf("json missing command_exited: %s", raw)
	}
}
