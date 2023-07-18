package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	fileSystem := os.DirFS(dirname)

	data, err := fs.ReadFile(fileSystem, ".aws/credentials")

	if err != nil {
		fmt.Printf("err: %v", err)
	}

	file_text := string(data)
	file_lines := strings.Split(file_text, "\n")

	access_key_index := 1
	secret_key_index := 2

	if len(args) == 0 {
		current_access_key := file_lines[access_key_index]
		for i, v := range file_lines {

			if v == "[default]" {
				continue
			}

			if strings.HasPrefix(v, "[") {
				if strings.HasPrefix(file_lines[i+1], current_access_key) {
					fmt.Print("=>")
				} else {
					fmt.Print("  ")
				}

				fmt.Println(v)
			}
		}

	} else {

		profile := args[0]

		match := fmt.Sprintf("[%v]", profile)

		for i, v := range file_lines {
			if v == match {
				access_key_index = i + 1
				secret_key_index = i + 2
			}
		}

		file_lines[1] = file_lines[access_key_index]
		file_lines[2] = file_lines[secret_key_index]

		new_file_text := strings.Join(file_lines, "\n")

		d1 := []byte(new_file_text)

		output_path := fmt.Sprintf("%v%v", dirname, "/.aws/credentials")

		err2 := os.WriteFile(output_path, d1, 0644)

		if err2 != nil {
			fmt.Printf("err: %v", err2)
		}

		fmt.Printf("Selected %v\n", match)
	}

}
