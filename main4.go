package main

import "fmt"

func main() {
	arrInputUser := []int{7, 11, 2, 15}

	targetInputUsr := 9

	var answerIndex []int
	for i := 0; i < len(arrInputUser); i++ {
		for k := i + 1; k < len(arrInputUser); k++ {
			if (arrInputUser[i] + arrInputUser[k]) == targetInputUsr {
				answerIndex = append(answerIndex, i)
				answerIndex = append(answerIndex, k)
			}
		}
	}

	fmt.Println(answerIndex)

}
