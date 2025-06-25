package fh3

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var contents []byte

func ReadAllEmbed() {
	fmt.Println("Contents of file:", string(contents))
}
