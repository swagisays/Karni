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
	fmt.Println("Connected to database")

	User := model.InitUserModel()

	data := map[string]interface{}{
		"email":    "   TEST125@TEST.com   ",
		"password": "12345678",
	}
	user := User.New(data)
	result, err := user.Save()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(user.Data)
}
