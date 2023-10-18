/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var checkCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the version of vcheck",
	Long: `Shows the current version of vcheck in the program.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1-alpha")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
