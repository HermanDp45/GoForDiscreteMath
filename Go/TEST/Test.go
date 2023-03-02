package main

import (
	//"errors"
	"fmt"
	//"unicode"
	"strconv"
)

func main() {
	//var users []string

	x, _ := strconv.ParseUint("123", 10, 64)
	fmt.Print(x)
}
