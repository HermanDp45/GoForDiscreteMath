package main

import (
    //"errors"
    "fmt"
    //"unicode"
    "strconv"

)

func main() {
    //var users []string

    //x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
    //fmt.Print(string('0'));
    x, _ := strconv.ParseUint("123", 10, 64)
    fmt.Print(x);
}