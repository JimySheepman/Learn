package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fileType := reflect.TypeOf((*os.File)(nil))

	fmt.Println(fileType.Implements(writerType))
}
