package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)

	str1 := `http://www.golangprograms.com/regular-expressions.html`
	match1 := re.FindStringSubmatch(str1)
	fmt.Println(match1[2])

	str2 := `/home/me/dir3/dir3a/dir3ac/filepat.png`
	match2 := re.FindStringSubmatch(str2)
	fmt.Println(match2[2])
}
