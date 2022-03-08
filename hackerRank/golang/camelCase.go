/*
Camel Case is a naming style common in many programming languages. In Java, method and variable names typically start with a lowercase letter, with all subsequent words starting with a capital letter (example: startThread). Names of classes follow the same pattern, except that they start with a capital letter (example: BlueCar).

Your task is to write a program that creates or splits Camel Case variable, method, and class names.

Input Format

    Each line of the input file will begin with an operation (S or C) followed by a semi-colon followed by M, C, or V followed by a semi-colon followed by the words you'll need to operate on.
    The operation will either be S (split) or C (combine)
    M indicates method, C indicates class, and V indicates variable
    In the case of a split operation, the words will be a camel case method, class or variable name that you need to split into a space-delimited list of words starting with a lowercase letter.
    In the case of a combine operation, the words will be a space-delimited list of words starting with lowercase letters that you need to combine into the appropriate camel case String. Methods should end with an empty set of parentheses to differentiate them from variable names.

Output Format

    For each input line, your program should print either the space-delimited list of words (in the case of a split operation) or the appropriate camel case string (in the case of a combine operation).

Sample Input

S;M;plasticCup()

C;V;mobile phone

C;C;coffee machine

S;C;LargeSoftwareBook

C;M;white sheet of paper

S;V;pictureFrame

Sample Output

plastic cup

mobilePhone

CoffeeMachine

large software book

whiteSheetOfPaper()

picture frame

Explanation

    Use Scanner to read in all information as if it were coming from the keyboard.

    Print all information to the console using standard output (System.out.print() or System.out.println()).

    Outputs must be exact (exact spaces and casing).
*/
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
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		worker(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("err")
	}
}
