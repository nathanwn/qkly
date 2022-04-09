package filesystem

import "testing"

func TestFileSystem(t *testing.T) {
	if Mgr() == nil {
		t.Fatalf("filesystem.Mgr() is not supposed to be nil\n")
	}
	if Mgr().Fs == nil {
		t.Fatalf("filesystem.Mgr().Fs is not supposed to be nil\n")
	}
	if Mgr().Util == nil {
		t.Fatalf("filesystem.Mgr().Util is not supposed to be nil\n")
	}
}
