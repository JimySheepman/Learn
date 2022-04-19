package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "4111111111111111"
	str2 := "346823285239073"
	str3 := "370750517718351"
	str4 := "4556229836495866"
	str5 := "5019717010103742"
	str6 := "76009244561"
	str7 := "4111-1111-1111-1111"
	str8 := "5610591081018250"
	str9 := "30569309025904"
	str10 := "6011111111111117"

	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nCC : %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("CC : %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("CC : %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("CC : %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("CC : %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("CC : %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("CC : %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("CC : %v :%v\n", str8, re.MatchString(str8))
	fmt.Printf("CC : %v :%v\n", str9, re.MatchString(str9))
	fmt.Printf("CC : %v :%v\n", str10, re.MatchString(str10))
}
