package tweet

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var id int64

const defaultTweetFormatTemplate = `
Created At: {{ .CreatedAt }}
Favorite Count: {{ .FavoriteCount }}
Retweet  Count: {{ .RetweetCount }}
Quote Count: {{ .QuoteCount }}
`

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the status of the Tweet",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		tweet, _, err := c.Statuses.Show(id, nil)

		if err != nil {
			log.Fatalf("could not show tweet: %v\n", err)
		}

		t := template.New("tweet_template")

		tmpl, err := t.Parse(defaultTweetFormatTemplate)
		if err != nil {
			log.Fatalf("failed to parse template: %v\n", err)
		}

		err = tmpl.Execute(os.Stdout, tweet)
		if err != nil {
			log.Fatalf("failed to execute template: %v\n", err)
		}
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "")
	},
}

func init() {
	TweetCmd.AddCommand(showCmd)
	showCmd.Flags().Int64VarP(&id, "id", "i", 0, "specify the id of the tweet")
}
