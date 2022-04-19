package account

import (
	"fmt"
	"log"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const UserInfoTemplate string = `Name: %s
Handle: @%s
Description: %s 
Followers Count: %d
Location: %s
`

// infoCmd represents the get command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get the detail account information",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		user, _, err := c.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
			IncludeEntities: twitter.Bool(true), SkipStatus: twitter.Bool(false), IncludeEmail: twitter.Bool(true),
		})

		if err != nil {
			log.Fatalf("could not get account info: %v\n", err)
		}

		fmt.Printf(UserInfoTemplate, user.Name,
			user.ScreenName,
			user.Description,
			user.FollowersCount,
			user.Location)
	},
}

func init() {
}
