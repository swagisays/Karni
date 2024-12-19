package model

import (
	karni "github.com/swagisays/karni/lib"
)

var userSchema = karni.Schema(map[string]karni.Field{
	"email": {Type: karni.String, Required: true, Trim: true, Lowercase: true, Unique: true},
	"password": {
		Type:     karni.String,
		Required: true, Trim: true, Lowercase: true,
	},
	"date": {
		Type:     karni.Date,
		Required: true,
	},
})

func InitUserModel() *karni.ModelStruct {
	User := karni.Model("users", userSchema)
	return User
}
