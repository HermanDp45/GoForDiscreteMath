package main

import (
	"fmt"
	. "fmt"
)

func FindEndBlock(str string, ind int) int {
	counter := 0
	for i:= ind; i < len(str); i++ {
		if (string(str[i]) == ")") && (counter == 0) {
			return i
		} else if (string(str[i]) == ")") && (counter != 0) {
			counter--
		} else if (string(str[i]) == "(") && (counter == 0) {
			counter++
		}
	}
	return -1
}

func CheckerExist(str string, strList []string) bool {
	for i := 0; i < len(strList); i++ {
		if (str == strList[i]) {
			return true
		}
	}
	
	return false
}

func CounterEco(str string, ind, countMove int, stringList []string) int {
	switch {
	case ind >= len(str):
		return countMove
	case (string(str[ind]) == "#")  || (string(str[ind]) == "@") || (string(str[ind]) == "$"):
		indEnd := FindEndBlock(str, ind)
		check := CheckerExist(str[ind - 1: indEnd + 1], stringList)
		//Printf("ind = %d   indEnd =   %d   check =  %b    string[] = %s\n", ind, indEnd, check, str[ind -1 : (indEnd + 1)])
		if !check {
			stringList = append(stringList, str[ind - 1: (indEnd + 1)])
			return CounterEco(str, (1 + ind), (1 + countMove), stringList)
		} else {
			return CounterEco(str, (1 + ind), countMove, stringList)
		}
	default:
		return CounterEco(str, (1 + ind), countMove, stringList)
	}
}

func main(){
	var str string
	fmt.Scanln(&str)

	ans := CounterEco(str, 0, 0, make([]string, 0, 5))

	Println(ans)

}