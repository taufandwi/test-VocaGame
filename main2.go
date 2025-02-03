package main

import "fmt"

func main() {
	//strInput := os.Args[1]
	strInput := "kakap"
	strInput2 := "kapaK"

	var flagArray []int
	for i := 0; i < len(strInput); i++ {
		for k := 0; k < len(strInput2); k++ {
			if strInput[i] == strInput2[k] {
				if len(flagArray) == 0 {
					flagArray = append(flagArray, k)
				} else {
					flagIsneverExist := false
					for i := 0; i < len(flagArray); i++ {
						if flagArray[i] == k {
							flagIsneverExist = true
							break
						}
					}

					if !flagIsneverExist {
						flagArray = append(flagArray, k)
					}
				}
			}
		}
	}

	//fmt.Println(len(flagArray))
	if len(flagArray) == len(strInput) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
