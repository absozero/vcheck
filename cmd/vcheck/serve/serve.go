/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package serve

import (
	"fmt"
	"github.com/absozero/vcheck/cmd/vcheck"
	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var VersionCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve Vcheck API",
	Long:  `Serves a continually updating API that returns package versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currently Under construction")
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
