package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tbugai/vue-echo/utils"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Create a new webapp",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]

		fmt.Printf("Creating '%s' app...\n", appName)
		if _, err := os.Stat(appName); os.IsNotExist(err) {
			err = os.Mkdir(appName, os.ModePerm)
		}
		utils.Copy("template/", appName)

		fmt.Printf("Fetching dependencies...\n")
		yarn := exec.Command("yarn", "install")
		yarn.Dir = appName + "/frontend"
		yarn.Output()

		echo := exec.Command("go", "get")
		echo.Dir = appName + "/backend"
		echo.Output()

		fmt.Printf("\nDone.\n")
	},
}
