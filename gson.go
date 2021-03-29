package main

import (
	"encoding/json"
)
func Encode(data map[string]interface{})[]byte{
	res,err := json.Marshal(data)
	if err != nil{
		panic(err)
	}
	return res
}