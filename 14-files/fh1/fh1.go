package fh1

import (
	"fmt"
	"os"
)

func ReadAll() {
	contents, err := os.ReadFile("/home/radosav/go/src/learngo/14-files/fh1/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(contents))
}
