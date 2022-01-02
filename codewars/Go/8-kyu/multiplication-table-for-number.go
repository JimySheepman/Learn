package kata

import (
	"fmt"
	"strings"
)

func MultiTable(number int) string {
	superstring := ""
	for i := 1; i < 11; i++ {
		superstring += fmt.Sprintf("%d * %d = %d\n", i, number, number*i)
	}
	return strings.TrimRight(superstring, "\n")
}
