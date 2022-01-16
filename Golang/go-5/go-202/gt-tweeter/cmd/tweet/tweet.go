package tweet

import (
	"os"

	"github.com/spf13/cobra"
)

var TweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Helps you to send tweet and show the status of the tweet",
	Long: `Basically you can send tweets and based on the command result you can query the status of the tweet to see
how many times it is retweeted, quoted, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	TweetCmd.AddCommand(sendCmd)
}
