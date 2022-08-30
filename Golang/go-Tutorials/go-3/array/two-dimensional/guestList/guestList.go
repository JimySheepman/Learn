package guestList

import (
	"github.com/Pallinder/go-randomdata"
)

func generate() [10][2]string {
	guestList := [10][2]string{}
	for index, _ := range guestList {
		guestList[index] = roomGuests(index)
	}
	return guestList
}

func roomGuests(roomId int) [2]string {
	guests := [2]string{}
	guests[0] = randomdata.SillyName()
	guests[1] = randomdata.SillyName()
	return guests
}
