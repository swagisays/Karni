package main

import (
	"fmt"
	"time"

	karni "github.com/swagisays/karni/lib"
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
		"email":    "   TEST1255@TEST.com   ",
		"password": "12345678",
		"date":     time.Now(),
	}
	user := User.New(data)
	result, err := user.Save()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(user.Data)
}
