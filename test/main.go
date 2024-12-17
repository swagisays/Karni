package main

import (
	"fmt"

	"github.com/swagisays/karni/karni"
	"github.com/swagisays/karni/test/model"
)

func main() {
	err := karni.Connect("mongodb://localhost:27017", "karniDb")
	if err != nil {
		fmt.Println(err)
	}

	User := model.InitUserModel()

	data := map[string]interface{}{
		"email":    "test@test.com",
		"password": "12345678",
	}
	user := User.New(data)
	user.Save()
}
