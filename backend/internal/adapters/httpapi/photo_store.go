package httpapi

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var errPhotoTooLarge = errors.New("photo too large")

type photoStore interface {
	Save(ctx context.Context, source io.Reader, maxBytes int64) (photoID string, sizeBytes int64, err error)
}

type filePhotoStore struct {
	baseDir string
}

func newFilePhotoStore(baseDir string) *filePhotoStore {
	return &filePhotoStore{baseDir: filepath.Clean(baseDir)}
}

func (s *filePhotoStore) Save(ctx context.Context, source io.Reader, maxBytes int64) (string, int64, error) {
	if err := os.MkdirAll(s.baseDir, 0o750); err != nil {
		return "", 0, fmt.Errorf("create photo dir: %w", err)
	}
	root, err := os.OpenRoot(s.baseDir)
	if err != nil {
		return "", 0, fmt.Errorf("open photo root: %w", err)
	}
	defer func() {
		if closeErr := root.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close photo root: %w", closeErr)
		}
	}()

	photoID, err := newPhotoID()
	if err != nil {
		return "", 0, err
	}

	dst, err := root.OpenFile(photoID, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return "", 0, fmt.Errorf("open destination file: %w", err)
	}

	success := false
	defer func() {
		_ = dst.Close()
		if !success {
			_ = root.Remove(photoID)
		}
	}()

	limited := &io.LimitedReader{R: source, N: maxBytes + 1}
	writtenBytes, err := io.Copy(dst, limited)
	if err != nil {
		return "", 0, fmt.Errorf("write photo file: %w", err)
	}
	if writtenBytes > maxBytes {
		return "", 0, errPhotoTooLarge
	}

	select {
	case <-ctx.Done():
		return "", 0, ctx.Err()
	default:
	}

	success = true
	return photoID, writtenBytes, nil
}

func newPhotoID() (string, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes[:]), nil
}
