package main

import (
	"fmt"
	"regexp"
)

func main() {
	var emojiRx = regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	var str = emojiRx.ReplaceAllString("Thats a nice joke ðŸ˜†ðŸ˜†ðŸ˜† ðŸ˜›", `[e]`)
	fmt.Println(str)
}
