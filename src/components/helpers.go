package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GenerateRandomString generates a cryptographically secure random string.
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

// FileExists checks if a file exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

// DirectoryExists checks if a directory exists.
func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return info.IsDir()
}

// CreateDirectory creates a directory if it does not exist.
func CreateDirectory(path string) error {
	if !DirectoryExists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// SanitizeFilename removes invalid characters from a filename.
func SanitizeFilename(filename string) string {
	invalidChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		filename = strings.ReplaceAll(filename, char, "_")
	}
	return filename
}

// GetTimestamp returns the current timestamp in RFC3339 format.
func GetTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

// JoinPath joins path elements using the correct separator for the OS.
func JoinPath(elem ...string) string {
	return filepath.Join(elem...)
}