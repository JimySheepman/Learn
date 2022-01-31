package main

import "fmt"

func main() {
	// Unsigned 8-bit integer,0 to 255
	var a uint8 = 5
	// Unsigned 16-bit integer,0 to 65535
	var b uint16 = 256
	// Unsigned 32-bit integer,0 to 4294967295
	var c uint32 = 65536
	// Unsigned 64-bit integer, 0 to 18446744073709551615
	var d uint64 = 4294967296

	fmt.Println(a, b, c, d)

	// Singed 8-bit integer, -128 to 127
	var e int8 = 5
	// Signed 16-bit integer, -32768 to 32767
	var f int16 = 128
	// Signed 32-bit integer, -2147483648 to 2147483647
	var g int32 = 32768
	// Signed 64-bit integer, -9223372036854775808 to 9223372036854775807
	var h int64 = 2147483648

	fmt.Println(e, f, g, h, )

	// 32-bit floating number
	var i float32 = 32.52
	// 64-bit floating number
	var j float64 = .587878

	fmt.Println(i, j)

	// 64-bit complex number
	var k complex64 = 2i + 5
	// 128-bit complex number
	var l complex128 = 2123i + .5

	fmt.Println(k, l)

	// Same as uint8
	var m byte = 5
	// Same as int32
	var n rune = 58848
	// 32 or 64 bit
	var o uint = 2545
	// 32 or 64 bit
	var p int = 5878
	// Uint for pointers
	var q uintptr = 12312

	fmt.Println(m, n, o, p, q)

	// Boolean type, true or false
	var r bool = true
	// String type
	var s string = "hey"
	fmt.Println(r, s)
}
