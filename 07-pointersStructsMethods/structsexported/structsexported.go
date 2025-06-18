package structsexported

import (
	"fmt"
	"learngo/07-pointersStructsMethods/structsexported/computer"
)

func StructExported() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
	}

	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
