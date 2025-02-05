/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	filehandler "aws-profile/src"
	"fmt"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select [profile]",
	Short: "Choose your default AWS profile from the list of existing profiles.",
	Long: `
Choose your default AWS profile from the list of existing profiles.
The Access Key and Secrey Key of the selected profile will be set as the keys of the default profile.`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		file_lines, err := filehandler.ReadCredentials()

		if err != nil {
			panic(err)
		}

		profile := args[0]
		match := fmt.Sprintf("[%v]", profile)

		selected_profile_index := -1
		has_default := false

		for i, v := range file_lines {
			if v == "[default]" {
				has_default = true
				continue
			}

			if v == match {
				selected_profile_index = i
			}
		}

		if !has_default {
			panic("You need to have a default profile to use the tool. To create a default profile, simply run 'aws configure' without the --profile flag.")
		}

		if selected_profile_index == -1 {
			fmt.Printf("Error: the profile %v doesn't exist\n", match)
			return
		}

		selected_access_key_index := selected_profile_index + 1
		selected_secret_key_index := selected_profile_index + 2

		file_lines[1] = file_lines[selected_access_key_index]
		file_lines[2] = file_lines[selected_secret_key_index]

		err = filehandler.WriteCredentials(file_lines)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Selected %v\n", match)

	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
