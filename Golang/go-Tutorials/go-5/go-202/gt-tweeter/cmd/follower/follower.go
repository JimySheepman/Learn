package follower

import (
	"os"

	"github.com/spf13/cobra"
)

var FollowerCmd = &cobra.Command{
	Use:   "follower",
	Short: "Helps you to list and get followers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	FollowerCmd.AddCommand(getCmd)
	FollowerCmd.AddCommand(listCmd)
}
