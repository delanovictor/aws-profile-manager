/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	filehandler "aws-profile/src"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the profiles on your aws/credentials file",
	Long: `
List the profiles on your aws/credentials file.
The profile that is also set as the default will be marked as selected.
If the default profile is not present as a labeled profile, it will be shown as 'Not Specified'.
`,
	Run: func(cmd *cobra.Command, args []string) {

		file_lines, err := filehandler.ReadCredentials()

		if err != nil {
			panic(err)
		}

		ak_index := 1

		current_ak := file_lines[ak_index]

		selected_profile := "[default]"
		has_default := false

		for i, v := range file_lines {

			if v == "[default]" {
				has_default = true
				continue
			}

			if strings.HasPrefix(v, "[") {
				if strings.HasPrefix(file_lines[i+1], current_ak) {
					fmt.Print("=>")
					selected_profile = v
				} else {
					fmt.Print("  ")
				}

				fmt.Println(v)
			}
		}

		if !has_default {
			panic("You need to have a default profile to use the tool. To create a default profile, simply run 'aws configure' without the --profile flag.")
		}

		ak_snippet := strings.ReplaceAll(current_ak, "aws_access_key_id=", "")

		ak_len := len(ak_snippet)

		ak_snippet = fmt.Sprintf("%v****************%v", ak_snippet[0:3], ak_snippet[ak_len-2:ak_len])
		fmt.Printf("\nCurrent Profile: %v - %v\n\n", selected_profile, ak_snippet)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
