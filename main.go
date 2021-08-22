package main

import (
	"fmt"

	"github.com/matishsiao/goInfo"
)

func main() {
	fmt.Println("Your OS")
	fmt.Println("__________")
	gi := goInfo.GetInfo()
	gi.VarDump()
}
