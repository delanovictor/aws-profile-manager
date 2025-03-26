package filehandler

import (
	"os"
	"bufio"
    "path/filepath"
	"strings"
	"runtime"
)

func ReadCredentials() ([]string, error) {
	user_directory, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

    credentials_path := filepath.Join(user_directory, ".aws/credentials")

	file, err := os.Open(credentials_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	var file_lines []string

	for s.Scan() {
		line := s.Text()

		file_lines = append(file_lines, line)
	}

	return file_lines, nil
}

func WriteCredentials(file_lines []string) error {
	user_directory, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	var file_text string
	
	if runtime.GOOS == "windows" {
		file_text = strings.Join(file_lines, "\r\n")
	} else {
		file_text = strings.Join(file_lines, "\n")
	}

	data := []byte(file_text)

    credentials_path := filepath.Join(user_directory, ".aws/credentials")

	return os.WriteFile(credentials_path, data, 0644)
}
