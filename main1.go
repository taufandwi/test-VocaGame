package main

import "fmt"

func main() {
	inptUser := "aaa"

	var strAnswer, currHuruf string
	counter := 0

	for i := 0; i < len(inptUser); i++ {
		if i == 0 {
			strAnswer = fmt.Sprintf("%c", inptUser[i])
			currHuruf = fmt.Sprintf("%c", inptUser[i])
			counter++
		} else if currHuruf == string(inptUser[i]) {
			counter++
		} else if currHuruf != string(inptUser[i]) {
			strAnswer += fmt.Sprintf("%d", counter)
			strAnswer += fmt.Sprintf("%c", inptUser[i])
			counter = 1
			currHuruf = fmt.Sprintf("%c", inptUser[i])
		}
	}
	strAnswer += fmt.Sprintf("%d", counter)

	if len(strAnswer) > len(inptUser) {
		fmt.Println(inptUser)
	} else {
		fmt.Println(strAnswer)
	}
}
