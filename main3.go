package main

import "fmt"

func main() {

	inptUsr := 100
	jmlTemp := 0

	currAngka := 1
	for i := 2; i < inptUsr; i++ {
		if i == 2 || isPrima(i) {
			jmlTemp += i
		}
		currAngka++
	}

	fmt.Println("Total :: ", jmlTemp)

}

func isPrima(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 && n != i {
			return false
		}
	}

	return true
}
