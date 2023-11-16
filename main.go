package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	arr := [...]string{"a", "b", "f"}
//encode
	bs, err := json.Marshal(arr)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

	//decode
	var name interface{}
	json.Unmarshal(bs, &name)

	fmt.Println(name)

}