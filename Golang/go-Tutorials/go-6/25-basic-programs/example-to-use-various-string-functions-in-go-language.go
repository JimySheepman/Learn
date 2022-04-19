package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-------------------------------")
	fmt.Println("Contain Function: ", strings.Contains("Australia", "lia"))
	fmt.Println("Contain Function: ", strings.Contains("Australia", "mia"))

	fmt.Println("-------------------------------")
	fmt.Println("Count Functions: ", strings.Contains("Australia", "lia"))
	fmt.Println("Count Functions: ", strings.Contains("Australia", "a"))
	fmt.Println("Count Functions: ", strings.Contains("Australia", "Bu"))

	fmt.Println("-------------------------------")
	fmt.Println("HasPrefix Function: ", strings.HasPrefix("Australia", "Bu"))
	fmt.Println("HasPrefix Function: ", strings.HasPrefix("Australia", "Aus"))

	fmt.Println("-------------------------------")
	fmt.Println("HasSuffix Function: ", strings.HasSuffix("Australia", "lia"))
	fmt.Println("HasSuffix Function: ", strings.HasSuffix("Australia", "tia"))

	fmt.Println("-------------------------------")
	fmt.Println("Index Function: ", strings.Index("Australia", "lia"))
	fmt.Println("Index Function: ", strings.Index("Australia", "a"))
	fmt.Println("Index Function: ", strings.Index("Australia", "A"))

	fmt.Println("-------------------------------")
	fmt.Println("Join Function: ", strings.Join([]string{"Australia", "Japan", "Germany"}, "-"))
	fmt.Println("Join Function: ", strings.Join([]string{"Australia", "Japan", "Germany"}, " "))

	fmt.Println("-------------------------------")
	fmt.Println("Repeat Function: ", strings.Repeat("Australia ", 3))

	fmt.Println("-------------------------------")
	fmt.Println("Replace Function: ", strings.Replace("Australia ", "a", "b", 1))
	fmt.Println("Replace Function: ", strings.Replace("Australia ", "a", "b", 2))
	fmt.Println("Replace Function: ", strings.Replace("Australia ", "a", "b", 3))
	fmt.Println("Replace Function: ", strings.Replace("Australia ", "A", "B", 1))
	fmt.Println("Replace Function: ", strings.Replace("Australia ", "lia", "pia", 1))

	fmt.Println("-------------------------------")
	fmt.Println("Split Function: ", strings.Split("Australia-Japan-Germany", "-"))

	fmt.Println("-------------------------------")
	fmt.Println("ToLower Function: ", strings.ToLower("Australia"))

	fmt.Println("-------------------------------")
	fmt.Println("ToUpper Function: ", strings.ToUpper("Australia"))
}
