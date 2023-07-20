package filehandler

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func ReadCredentials() ([]string, error) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	fileSystem := os.DirFS(dirname)

	data, err := fs.ReadFile(fileSystem, ".aws/credentials")

	if err != nil {
		return nil, err
	}

	file_text := string(data)

	file_text = strings.Replace(file_text, "\n\n", "\n", -1)

	file_lines := strings.Split(file_text, "\n")

	return file_lines, nil
}

func WriteCredentials(file_lines []string) error {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	file_text := strings.Join(file_lines, "\n")

	data := []byte(file_text)

	output_path := fmt.Sprintf("%v%v", dirname, "/.aws/credentials")

	return os.WriteFile(output_path, data, 0644)
}
