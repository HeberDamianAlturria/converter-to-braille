package output

import (
	"os"
	"strings"
	"errors"
)

// SaveTextFile saves the given content to a text file at the specified path.
func SaveTextFile(path, content string) error {
	if strings.TrimSpace(path) == "" {
		return errors.New("path cannot be empty")
	}

	if strings.TrimSpace(content) == "" {
		return errors.New("content cannot be empty")
	}

	if !strings.HasSuffix(path, ".txt") {
		return errors.New("file path must end with .txt")
	}

	// Create the file at the specified path.
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Write the content to the file.
	_, err = file.WriteString(content)

	return err
}