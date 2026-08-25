package ttywatch

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func TestReserveCustomSessionIDReclaiming_aliveClaim(t *testing.T) {
	dir := t.TempDir()
	cfg := DefaultRegistryConfig(dir)
	sid := "codex-status-usage"
	regDir := RegistryDir(cfg)
	if err := os.MkdirAll(regDir, 0755); err != nil {
		t.Fatal(err)
	}
	claim := filepath.Join(regDir, "."+sid+".claim")
	if err := os.WriteFile(claim, []byte(strconv.Itoa(os.Getpid())+"\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// Plain reserve must fail while the alive claim is present.
	if _, err := ReserveCustomSessionID(cfg, sid); err == nil || !isSessionIDAlreadyInUse(err) {
		t.Fatalf("plain reserve want already in use, got %v", err)
	}

	release, err := ReserveCustomSessionIDReclaiming(cfg, sid)
	if err != nil {
		t.Fatalf("Reclaiming: %v", err)
	}
	release()
	clearSessionClaim(cfg, sid)
}
