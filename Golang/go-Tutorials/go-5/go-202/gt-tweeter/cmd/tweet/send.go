package tweet

import (
	"fmt"
	"log"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var message string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a Tweet with the given message",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		tweet, _, err := c.Statuses.Update(message, nil)

		if err != nil {
			log.Fatalf("coult not send tweet: %v\n", err)
		}

		fmt.Printf(`Tweet created successfully
You can check the status of the Tweet by running the following command:
$ gt-tweeter tweet show --id %d`, tweet.ID)
	},
}

func init() {
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "specify the message")
}
