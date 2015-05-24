package main

import (
	"fmt"
	"ynote"
)

func main() {
	yClient := ynote.New()
	msg := recover()
	fmt.Print(msg)
	yClient.Login()

}
