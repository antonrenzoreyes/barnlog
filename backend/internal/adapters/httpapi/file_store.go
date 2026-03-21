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
	"strings"
)

var errFileTooLarge = errors.New("file too large")

type fileStore interface {
	Save(ctx context.Context, source io.Reader, maxBytes int64) (fileID string, sizeBytes int64, err error)
}

type diskFileStore struct {
	baseDir string
}

func newFileStore(baseDir string) *diskFileStore {
	trimmed := strings.TrimSpace(baseDir)
	cleaned := filepath.Clean(trimmed)
	if trimmed == "" || cleaned == "." || cleaned == string(filepath.Separator) {
		return nil
	}
	return &diskFileStore{baseDir: cleaned}
}

func (s *diskFileStore) Save(
	ctx context.Context,
	source io.Reader,
	maxBytes int64,
) (fileID string, sizeBytes int64, err error) {
	if err := os.MkdirAll(s.baseDir, 0o750); err != nil {
		return "", 0, fmt.Errorf("create file dir: %w", err)
	}
	root, err := os.OpenRoot(s.baseDir)
	if err != nil {
		return "", 0, fmt.Errorf("open file root: %w", err)
	}
	defer func() {
		if closeErr := root.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close file root: %w", closeErr)
		}
	}()

	fileID, err = newFileID()
	if err != nil {
		return "", 0, err
	}

	dst, err := root.OpenFile(fileID, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return "", 0, fmt.Errorf("open destination file: %w", err)
	}

	success := false
	defer func() {
		if closeErr := dst.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close destination file: %w", closeErr)
		}
		if !success {
			_ = root.Remove(fileID)
		}
	}()

	limited := &io.LimitedReader{R: source, N: maxBytes + 1}
	buffer := make([]byte, 32*1024)
	var writtenBytes int64
	for {
		select {
		case <-ctx.Done():
			return "", 0, ctx.Err()
		default:
		}

		readBytes, readErr := limited.Read(buffer)
		if readBytes > 0 {
			writtenNow, writeErr := dst.Write(buffer[:readBytes])
			writtenBytes += int64(writtenNow)
			if writeErr != nil {
				return "", 0, fmt.Errorf("write file: %w", writeErr)
			}
			if writtenNow != readBytes {
				return "", 0, fmt.Errorf("write file: %w", io.ErrShortWrite)
			}
			if writtenBytes > maxBytes {
				return "", 0, errFileTooLarge
			}
		}

		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				break
			}
			return "", 0, fmt.Errorf("write file: %w", readErr)
		}
	}

	success = true
	return fileID, writtenBytes, nil
}

func newFileID() (string, error) {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes[:]), nil
}
