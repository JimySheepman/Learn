package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Options struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func New(co Options) *twitter.Client {
	config := oauth1.NewConfig(co.ConsumerKey, co.ConsumerSecret)
	token := oauth1.NewToken(co.AccessToken, co.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}
