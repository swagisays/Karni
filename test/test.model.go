package main

import (
	"fmt"

	"github.com/swagisays/karni/karni"
)

var userSchema = karni.Schema(map[string]karni.Field{
	"email": {Type: karni.String, Required: true, Lowercase: true},
	"password": {
		Type: karni.String,
		Validators: []func(interface{}) error{
			func(value interface{}) error {
				str := value.(string)
				if len(str) < 8 {
					return fmt.Errorf("password too short")
				}
				return nil
			},
		},
	},
})

var User = karni.Model("users", userSchema)
