package main

import (
	"fmt"

	"github.com/swagisays/karni/karni"
)

func main() {
	err := karni.Connect("mongodb://localhost:27017", "akjdhaskjdh")
	if err != nil {
		fmt.Println(err)
	}
}
