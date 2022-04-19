package follower

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/dghubble/go-twitter/twitter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var count int

const defaultFollowerFormatTemplate = `Follower: {{ .Name }}@{{ .ID }}, following: {{ .Following }}
👉 https://twitter.com/@{{ .ScreenName }}`

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the followers",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		s, _ := cmd.Flags().GetInt("count")

		followers, _, err := c.Followers.List(&twitter.FollowerListParams{
			Count: s,
		})

		if err != nil {
			log.Fatalf("could not list followers: %v\n", err)
		}

		t := template.New("follower_template")

		for index, follower := range followers.Users {
			tmpl, err := t.Parse(defaultFollowerFormatTemplate)
			if err != nil {
				log.Fatalf("failed to parse template: %v\n", err)
			}

			fmt.Fprintf(os.Stdout, "%d. ", index)
			err = tmpl.Execute(os.Stdout, follower)
			if err != nil {
				log.Fatalf("failed to execute template: %v\n", err)
			}
			fmt.Fprintln(os.Stdout, "")
			fmt.Fprintln(os.Stdout, "")
		}
	},
}

func init() {
	listCmd.Flags().IntVarP(&count, "count", "c", 20, "specify the count of the followers")
}
