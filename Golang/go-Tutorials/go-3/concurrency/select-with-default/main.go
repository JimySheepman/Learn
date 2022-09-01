package main

import "log"

func main() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	select {
	case rec, ok := <-ch1:
		if ok {
			log.Printf("received on ch1 : %s", rec)
		}
	case rec, ok := <-ch2:
		if ok {
			log.Printf("received on ch2 : %s", rec)
		}
	default:
		log.Println("default case")
	}
	log.Println("end")
}
