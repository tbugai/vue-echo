package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of vue-echo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Vue.js/Echo Webapp Tool v0.1 -- HEAD")
	},
}
