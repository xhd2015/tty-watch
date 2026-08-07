package ttywatch

import (
	"fmt"
	"testing"
)

func TestIsSessionIDAlreadyInUse(t *testing.T) {
	if !isSessionIDAlreadyInUse(fmt.Errorf(`run: session id "x" already in use`)) {
		t.Fatal("expected match")
	}
	if isSessionIDAlreadyInUse(fmt.Errorf("other")) {
		t.Fatal("expected no match")
	}
	if isSessionIDAlreadyInUse(nil) {
		t.Fatal("nil")
	}
}

func TestShouldReclaimZombieForReserve_commandExited(t *testing.T) {
	dir := t.TempDir()
	cfg := RegistryConfig{Home: dir, Subdir: "codex-tty-registry"}
	sid := "sess-z-1"
	if err := WriteRegistry(cfg, RegistryEntry{
		SessionID:     sid,
		ListenAddr:    "127.0.0.1:9",
		PID:           1,
		Command:       []string{"codex"},
		CommandPID:    2,
		CommandExited: true,
	}); err != nil {
		t.Fatal(err)
	}
	if !shouldReclaimZombieForReserve(cfg, sid) {
		t.Fatal("command_exited should reclaim")
	}
}

func TestShouldReclaimZombieForReserve_missingEntry(t *testing.T) {
	dir := t.TempDir()
	cfg := RegistryConfig{Home: dir, Subdir: "codex-tty-registry"}
	if !shouldReclaimZombieForReserve(cfg, "no-such") {
		t.Fatal("missing entry should reclaim (claim-only hold)")
	}
}
