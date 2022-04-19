package follower

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var id int64

const defaultFollowerDetailFormatTemplate = `Follower: {{ .Name }}
Following: {{ .Following }}
Verified: {{ .Verified }}
URL: {{ .URL }}
Followers Count: {{ .FollowersCount }}
Friends Count: {{ .FriendsCount }}
Location: {{ .Location }}
-> https://twitter.com/@{{ .ScreenName }}`

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the follower with the given id",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		id, _ := cmd.Flags().GetInt64("id")

		follower, _, err := c.Followers.List(&twitter.FollowerListParams{
			UserID: id,
		})

		if err != nil {
			log.Fatalf("could not get follower: %v\n", err)
		}

		t := template.New("follower_template")

		tmpl, err := t.Parse(defaultFollowerDetailFormatTemplate)
		if err != nil {
			log.Fatalf("failed to parse template: %v\n", err)
		}

		err = tmpl.Execute(os.Stdout, follower.Users[0])
		if err != nil {
			log.Fatalf("failed to execute template: %v\n", err)
		}
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "")

	},
}

func init() {
	getCmd.Flags().Int64VarP(&id, "id", "i", 0, "specify the id of the follower")
}
