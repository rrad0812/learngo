package structsexported

import (
	"fmt"
	"learngo/16-structs/structsexported/computer"
)

func StructExported() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}

	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
