package main

import "testing"

type Quote string
type Movie string

const (
	Crush  Quote = "To cruch your enemies..."
	T1000  Quote = "I'll be back"
	Unknow Quote = "unknow"

	Conan       Movie = "conan"
	Terminator2 Movie = "terminator2"
	Predator    Movie = "predator"
)

func MovieQuote(movie Movie) Quote {
	switch movie {
	case Conan:
		return Cruch
	case Terminator2:
		return T1000
	default:
		return Unknow
	}
}

func TestMovieQuote(t *testing.T) {
	movies := []Movie{Conan, Predator, Terminator2}
	for _, m := range movies {
		if q := MovieQuote(m); q == Unknow {
			t.Error("unknow quote for movie", m)
		}
	}
}
