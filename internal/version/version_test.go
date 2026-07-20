package version

import (
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	info := Get()

	if info.Version == "" {
		t.Fatal("version is empty")
	}

	if info.Commit == "" {
		t.Fatal("commit is empty")
	}

	if info.BuildDate == "" {
		t.Fatal("build date is empty")
	}

	if info.GoVersion != runtime.Version() {
		t.Fatalf("unexpected Go version: %s", info.GoVersion)
	}
}

func TestString(t *testing.T) {
	s := Get().String()

	if s == "" {
		t.Fatal("string output is empty")
	}
}
