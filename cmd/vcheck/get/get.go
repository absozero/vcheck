/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package version

import (
	"github.com/absozero/vcheck/cmd/vcheck"
	"github.com/absozero/vcheck/pkg/vercheck"
	"github.com/spf13/cobra"
	"fmt"
)

// helpCmd represents the help command
var VersionCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns the latest version of a given project on github",
	Long: `Shows the current version of vcheck in the program.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(vercheck.GithubCall())
	},
}

func init() {

	cmd.VcheckCmd.AddCommand(VersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
