package main

import (
	"errors"
	"strings"
	"time"
)

func GetMockStream() Stream {
	return Stream{
		pos:    0,
		tweets: mockdata,
	}
}

type Stream struct {
	pos    int
	tweets []Tweet
}

var ErrEOF = errors.New("end of file")

func (s *Stream) Next() (*Tweet, error) {
	time.Sleep(320 * time.Millisecond)
	if s.pos >= len(s.tweets) {
		return &Tweet{}, ErrEOF
	}

	tweet := s.tweets[s.pos]
	s.pos++

	return &tweet, nil
}

type Tweet struct {
	Username string
	Text     string
}

func (t *Tweet) IsTalkingAboutGo() bool {
	time.Sleep(330 * time.Millisecond)

	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher

}

var mockdata = []Tweet{
	{
		Username: "davecheney",
		Text:     "#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		Username: "beertocode",
		Text:     "Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		Username: "ironzeb",
		Text:     "Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		Username: "beertocode",
		Text:     "Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		Username: "vampirewalk666",
		Text:     "I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}
