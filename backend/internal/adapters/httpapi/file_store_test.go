package httpapi

import (
	"context"
	"errors"
	"io"
	"testing"
)

func TestNewFileStoreRejectsInvalidBaseDir(t *testing.T) {
	t.Parallel()

	invalidDirs := []string{"", "   ", ".", "/"}
	for _, dir := range invalidDirs {
		if store := newFileStore(dir); store != nil {
			t.Fatalf("expected nil store for invalid dir %q", dir)
		}
	}
}

func TestNewFileStoreAcceptsValidBaseDir(t *testing.T) {
	t.Parallel()

	store := newFileStore("backend/uploads/files")
	if store == nil {
		t.Fatalf("expected non-nil store for valid dir")
	}
}

func TestFileStoreSaveHonorsContextCancellation(t *testing.T) {
	t.Parallel()

	storeDir := t.TempDir()
	store := newFileStore(storeDir)
	if store == nil {
		t.Fatalf("expected non-nil store for temp dir")
	}

	ctx, cancel := context.WithCancel(context.Background())
	source := &cancelAfterFirstRead{cancel: cancel}

	_, _, err := store.Save(ctx, source, 1024)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}
	if source.secondRead {
		t.Fatalf("expected save to stop reading once context is canceled")
	}
}

type cancelAfterFirstRead struct {
	cancel     context.CancelFunc
	read       bool
	secondRead bool
}

func (r *cancelAfterFirstRead) Read(p []byte) (int, error) {
	if !r.read {
		r.read = true
		copy(p, "abc")
		r.cancel()
		return 3, nil
	}

	r.secondRead = true
	copy(p, "def")
	return 3, io.EOF
}
