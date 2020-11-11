package filesys

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Abs ...
func Abs(src string) (string, error) {
	path, err := filepath.Abs(src)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(path, "\\", "/"), err
}

// SplitSlugs ...
func SplitSlugs(path string) []string {
	return strings.Split(strings.Trim(strings.ReplaceAll(path, "\\", "/"), "/"), "/")
}

// JoinSlugs ... TODO: write tests
func JoinSlugs(slugs ...string) string {
	return strings.Trim(strings.Join(slugs, "/"), "/")
}

// CreateFile ... TODO: write tests and make it create directories too
func CreateFile(path string) (*os.File, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), 0755)
	}

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return os.Create(path)
}

// CommonSuffix ...
func CommonSuffix(dst, src string) string {
	// First we need to find out where exactly these two paths start to differ, so we need to iterate over both paths in reverse
	dl := len(dst)
	sl := len(src)

	for i := 0; i < dl; i++ {

		// Outside the range for the source
		if i >= sl {
			return dst[dl-i:]
		}

		// Find the first character that doesn't match
		if src[sl-i-1] != dst[dl-i-1] {
			return dst[dl-i:]
		}
	}
	return dst
}

// Checksum a file
func Checksum(path string, hash hash.Hash) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
