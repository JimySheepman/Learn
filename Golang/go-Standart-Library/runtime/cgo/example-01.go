package main

/*
extern void MyGoPrint(void *context);
static inline void myprint(void *context) {
    MyGoPrint(context);
}
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export MyGoPrint
func MyGoPrint(context unsafe.Pointer) {
	h := *(*cgo.Handle)(context)
	val := h.Value().(string)
	println(val)
	h.Delete()
}

func main() {
	val := "hello Go"
	h := cgo.NewHandle(val)
	C.myprint(unsafe.Pointer(&h))
	// Output: hello Go
}
