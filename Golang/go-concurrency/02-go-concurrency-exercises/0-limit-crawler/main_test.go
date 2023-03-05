package main

import (
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	fetchSig := fetchSignalInstance()

	start := time.Unix(0, 0)
	go func(start time.Time) {
		for {
			switch {
			case <-fetchSig:
				if time.Now().Sub(start).Nanoseconds() < 950000000 {
					t.Log("There exists a two crawls that were executed less than 1 second apart.")
					t.Log("Solution is incorrect.")
					t.FailNow()
				}
				start = time.Now()
			}
		}
	}(start)

	main()
}
