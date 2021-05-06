package main

import (
	"fmt"
	"github.com/LouisBillaut/capMonsterTool"
)

//this example get your actual balance account

const apiKey = "67b6bcbb1a728ea8d563de6d169a2057"

func main() {
	res, err := capMonsterTool.GetBalance(apiKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("My balance:", res)
}