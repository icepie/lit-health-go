package main

import (
	"fmt"
	"lit-healthy-go/lithe"
)

func main() {
	lgrt := lithe.HealthyLogin("CarddNO", "Password") //set the return value of login
	fmt.Println(lgrt)
}
