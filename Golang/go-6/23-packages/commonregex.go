package main

import (
	"fmt"

	cregex "github.com/mingrammer/commonregex"
)

func main() {
	text := `John, please get that article on www.linkedin.com 
			to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, actually. 
			If you have any questions, You can reach me at (519)-236-2723x341 or 234-567-8900 or +41 22 730 5989
			or get in touch with my associate at harold.smith@gmail.com. You system details as below:
			fe80:0:0:0:204:61ff:fe9d:f156, 192.30.253.113`

	dateList := cregex.Date(text)
	fmt.Println(dateList)

	timeList := cregex.Time(text)
	fmt.Println(timeList)

	linkList := cregex.Links(text)
	fmt.Println(linkList)

	ipList := cregex.IPs(text)
	fmt.Println(ipList)

	IPv4List := cregex.IPv4s(text)
	fmt.Println(IPv4List)

	IPv6List := cregex.IPv6s(text)
	fmt.Println(IPv6List)

	emailList := cregex.Emails(text)
	fmt.Println(emailList)

	phones := cregex.Phones(text)
	fmt.Println(phones)

	phoneList := cregex.PhonesWithExts(text)
	fmt.Println(phoneList)

	text = "price is $1,000. Pay using Credit card 4111 1111 1111 1111 and address is 504 parkwood drive, 02540, US"
	creditCard := cregex.CreditCards(text)
	fmt.Println(creditCard)
	price := cregex.Prices(text)
	fmt.Println(price)

	address := "504 parkwood drive, 02540, US"
	zip := cregex.ZipCodes(address)
	fmt.Println(zip)
	streetAddress := cregex.StreetAddresses(address)
	fmt.Println(streetAddress)
}
