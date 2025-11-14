package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func split(str string) string {
	re := regexp.MustCompile(`([A-Z]+)`)
	str = re.ReplaceAllString(str, ` $1`)
	str = strings.Trim(str, " ")
	return str
}

func combine(str string) string {
	s := strings.Split(str, " ")
	newS := ""
	for _, v := range s {
		if newS == "" {
			newS = v
		} else {
			newS += strings.Title(v)
		}
	}
	return newS
}

func worker(str string) {
	command := strings.Split(str, ";")
	if command[0] == "S" {
		newStr := split(command[2])
		switch command[1] {
		case "M":
			newStr1 := strings.Split(newStr, "()")
			s := ""
			for _, v := range newStr1 {
				s += (v + " ")
			}
			fmt.Println(strings.ToLower(s))
			break
		case "C":
			fmt.Println(strings.ToLower(newStr))
			break
		case "V":
			fmt.Println(strings.ToLower(newStr))
			break
		}
	} else if command[0] == "C" {
		newStr := combine(command[2])
		switch command[1] {
		case "M":
			fmt.Println(newStr + "()")
			break
		case "C":
			fmt.Println(strings.Title(newStr))
			break
		case "V":
			fmt.Println(newStr)
			break
		}
	}
}
func main() {
	//var input string
	//fmt.Scanln(&input)
	//worker(input)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		worker(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("err")
	}
}
