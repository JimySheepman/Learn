package account

import (
	"os"

	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Helps you to view account information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	AccountCmd.AddCommand(infoCmd)
}
