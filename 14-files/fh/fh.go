package fh

import (
	"fmt"
	"os"
)

func ReadAll() {
	contents, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
