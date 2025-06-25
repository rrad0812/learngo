package fh2

import (
	"flag"
	"fmt"
	"os"
)

func ReadAllFlag() {

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	fmt.Println("value of fpath is", *fptr)

	contents, err := os.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))

}
