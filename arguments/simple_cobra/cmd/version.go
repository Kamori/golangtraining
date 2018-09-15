package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var text string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of example",
	Long:  `All software has versions. This is the examples's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1 Beta", text)
	},
}

func Execute() {
	versionCmd.Flags().StringVar(&text, "extratext", "", "Extra text to append to version string")

	if err := versionCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
