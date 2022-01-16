package cmd

import (
	"fmt"
	"os"

	"github.com/GoTurkiye/training/104-cli/gt-tweeter/cmd/account"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/cmd/completion"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/cmd/follower"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/cmd/tweet"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gt-tweeter",
	Short: "Facilitate to use Twitter from your terminal",
	Long: `gt-tweeter is a CLI library for Go that facilitates to use Twitter from your terminal.
It basically shows account information, list and get followers, or event you can send tweet, and
display the status of the tweet.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gt-tweeter.yaml)")
	rootCmd.AddCommand(follower.FollowerCmd)
	rootCmd.AddCommand(account.AccountCmd)
	rootCmd.AddCommand(tweet.TweetCmd)
	rootCmd.AddCommand(completion.CompletionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".gt-tweeter")
		viper.SetConfigType("yaml")
		viper.SetConfigType("yml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
