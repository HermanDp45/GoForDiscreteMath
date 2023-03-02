package main

import (
	"fmt"
	"os"
	"io"
)

func main() {


	str := ""


	file, err := os.Open("input.txt")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
     
    data := make([]byte, 64)
     
    for{
        n, err := file.Read(data)
        if err == io.EOF{  
            break          
        }
        str += string(data[:n])
    }

	
	counts := make([]int, 256)
	countNotZero := 0
	strout := ""
	maxc := 0
	for _, char := range str {
		counts[char]++
		if counts[char] > maxc && (char > 32) && (char < 123){
			maxc++
		}
		if counts[char] == 1 {
			countNotZero++
		}
	}

	for i := 0; i < 256; i++ {
		if counts[i] > 0  && (i > 32) && (i < 123) {
			strout += string(i)
		}
	}

	for ; maxc > 0; maxc-- {
		for _, char := range strout {
			if counts[char] == maxc {
				fmt.Print("#")
				counts[char]--
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print(strout)
}
