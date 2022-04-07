package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	bReader := bytes.NewReader([]byte("some secret text"))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", out.Bytes())
}
