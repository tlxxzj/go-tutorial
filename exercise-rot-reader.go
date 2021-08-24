package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rotate13(char byte) byte {
	if char >= 'a' && char <= 'z' {
		if char <= 'm' {
			return char + 13
		} else {
			return char - 13
		}
	} else if char >= 'A' && char <= 'Z' {
		if char <= 'M' {
			return char + 13
		} else {
			return char - 13
		}
	}
	return char
}

func (rot13 rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot13.r.Read(bytes)
	for i := 0; i < n; i++ {
		bytes[i] = Rotate13(bytes[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
