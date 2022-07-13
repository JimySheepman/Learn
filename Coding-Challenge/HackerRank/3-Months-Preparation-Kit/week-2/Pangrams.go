package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	str := strings.ToLower(s)
	splitAlphabase := strings.Split(alphabase, "")
	splitStr := strings.Split(str, "")
	dict := make(map[string]bool, len(splitAlphabase))
	for _, v := range splitAlphabase {
		dict[v] = false
	}
	for _, v := range splitStr {
		if v != " " {
			dict[v] = true
		}
	}
	for _, v := range dict {
		if v == false {
			return "not pangram"
		}
	}

	return "pangram"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
