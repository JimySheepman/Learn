package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `<html><body>
			<form name="query" action="http://www.example.net/action.php" method="post">
				<textarea type="text" name="nameiknow">The text I want</textarea>
				<div id="button">
					<input type="submit" value="Submit" />
				</div>
			</form>
			</body></html>`

	re := regexp.MustCompile(`<textarea.*?>(.*)</textarea>`)

	submatchall := re.FindAllStringSubmatch(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}
}
