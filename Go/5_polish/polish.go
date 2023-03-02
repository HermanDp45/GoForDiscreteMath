package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func FindCloseInd(str string) int {
	counter := 0
	for i := 0;; i++ {
		if string(str[i]) == "(" {
			counter++
		} else if (string(str[i]) == ")") && (counter == 0) {
			return i
		} else if (string(str[i]) == ")") && (counter != 0) {
			counter--
		}
	}
}

func RecCounter(str string, ind int) int {
	switch {
	case string(str[ind]) == "*":
		retIndR := FindCloseInd(str[ind:]) + 1
		if string(str[ind + 2]) == "("{
			retIndL := FindCloseInd(str[ind + 3:]) + 1
			return RecCounter(str[ind + 2 : ind + 2 + retIndL + 1], 0) * RecCounter(str[ind + retIndL + 4 : retIndR], 0)
		} else {
			return RecCounter(str, ind + 2) * RecCounter(str[ind + 4 : retIndR], 0)
		}

	case string(str[ind]) == "+":
		retIndR := FindCloseInd(str[ind:]) + 1
		if string(str[ind + 2]) == "("{
			retIndL := FindCloseInd(str[ind + 3:]) + 1
			return RecCounter(str[ind + 2 : ind + 2 + retIndL + 1], 0) + RecCounter(str[ind + retIndL + 4 : retIndR], 0)
		} else {
			return RecCounter(str, ind + 2) + RecCounter(str[ind + 4 : retIndR], 0)
		}

	case string(str[ind]) == "-":
		retIndR := FindCloseInd(str[ind:]) + 1
		if string(str[ind + 2]) == "("{
			retIndL := FindCloseInd(str[ind + 3:]) + 1
			return RecCounter(str[ind + 2 : ind + 2 + retIndL + 1], 0) - RecCounter(str[ind + retIndL + 4 : retIndR], 0)
		} else {
			return RecCounter(str, ind + 2) - RecCounter(str[ind + 4 : retIndR], 0)
		}

	case strings.Contains("0123456789", string(str[ind])):
		number, _ := strconv.Atoi(string(str[ind]))
		return number 

	default:
		return RecCounter(str[1:], 0)
	}
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text()
	Println(RecCounter(str, 0))
}