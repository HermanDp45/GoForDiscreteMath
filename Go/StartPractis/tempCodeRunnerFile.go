package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str := ""
	for {
		line, err:= reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		str += line
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
