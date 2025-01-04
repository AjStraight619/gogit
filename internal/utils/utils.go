package utils

import (
	"crypto"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func ResolvePath(path string) (string, error) {
	if path == "." {
		cwd, err := os.Getwd()

		if err != nil {
			return "", fmt.Errorf("failed to get current working directory: %w", err)
		}

		path = cwd
	}

	absPath, err := filepath.Abs(path)

	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for path: %s, %w", path, err)
	}

	return absPath, nil

}

func statFile(path string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("path does not exist: %s", path)
		}
		if os.IsPermission(err) {
			return nil, fmt.Errorf("permission denied: %s", path)
		}
		// Catch-all for other errors
		return nil, fmt.Errorf("failed to stat path %s: %w", path, err)
	}

	return fileInfo, nil

}

func ProcessPaths(path string, filesInfo *[]os.FileInfo) error {
	fmt.Printf("Processing path: %s\n", path)

	// Split the path into root directory and relative part
	root := filepath.Dir(path)
	relative := filepath.Base(path)

	// Use WalkDir to process the directory tree
	return fs.WalkDir(os.DirFS(root), relative, func(entryPath string, d fs.DirEntry, err error) error {
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Path does not exist during walk: %s\n", entryPath)
				return nil // Log and continue
			}
			if os.IsPermission(err) {
				fmt.Printf("Permission denied during walk: %s\n", entryPath)
				return nil // Log and continue
			}
			// Catch-all for other errors
			return fmt.Errorf("error accessing path '%s': %w", entryPath, err)
		}

		// Build the absolute path
		absolutePath := filepath.Join(root, entryPath)

		// Stat the current entry
		fileInfo, err := os.Stat(absolutePath)
		if err != nil {
			return fmt.Errorf("failed to stat path '%s': %w", absolutePath, err)
		}

		// Log and add to filesInfo
		if d.IsDir() {
			fmt.Printf("Directory: %s\n", absolutePath)
		} else {
			fmt.Printf("File: %s\n", absolutePath)
			*filesInfo = append(*filesInfo, fileInfo)
		}
		return nil
	})
}

func CalculateChecksum(path string) (string, error) {

	file, err := os.Open(path)

	if err != nil {
		return "", fmt.Errorf("Failed to calculate checksum for path: %s", path)
	}

	defer file.Close()

	hasher := crypto.BLAKE2b_256.New()

	buf := make([]byte, 4096)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("failed to read file: %v", err)
		}

		hasher.Write(buf[:n])

	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
