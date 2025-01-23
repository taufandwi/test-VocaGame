package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Car struct {
	Number     string
	IsAllocate bool
}

func main() {
	var totalParkingSize int64

	parkingCounter := 1

	listParkedCars := make(map[string]*Car)

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}
	// data is the file content, you can use it
	fmt.Println("File content is:")

	// break into each command
	readDataFile := string(data)
	listCommand := strings.Split(readDataFile, "\n")

	for _, line := range listCommand {
		//fmt.Println(line)
		commands := strings.Split(line, " ")
		if len(commands) > 0 {
			switch strings.TrimSpace(commands[0]) {
			case "create_parking_lot":
				totalParkingSize, err = strconv.ParseInt(strings.TrimSpace(commands[1]), 10, 64)
				if err != nil {
					fmt.Println("Parking size is invalid", err)
				}

				break
			case "park":
				if totalParkingSize > int64(len(listParkedCars)) && totalParkingSize >= int64(parkingCounter) {
					listParkedCars[strconv.Itoa(parkingCounter)] = &Car{strings.TrimSpace(commands[1]), true}
					fmt.Println("Allocated slot number: ", parkingCounter)
					parkingCounter++
				} else if totalParkingSize >= int64(parkingCounter) {
					for parkID, parkingLot := range listParkedCars {
						if parkingLot.IsAllocate == false {
							listParkedCars[parkID] = &Car{strings.TrimSpace(commands[1]), true}
							fmt.Println("Allocated slot number: ", parkID)
							break
						}
					}
				} else {
					fmt.Println("Sorry, parking lot is full")
				}

				fmt.Println("")
				break
			case "leave":
				for parkID, car := range listParkedCars {
					if car.Number == commands[1] {
						car.IsAllocate = false
						durrInt, err := strconv.Atoi(strings.TrimSpace(commands[2]))
						if err != nil {
							fmt.Println("Duration is invalid", err)
						}
						fmt.Printf("Registration number %s with Slot Number %s is free with Charge $%v\n", car.Number, parkID, getTotalCharges(durrInt))
						parkingCounter--
						break
					}
				}

				fmt.Println("")
				break
			case "status":
				printStatus(listParkedCars)

				fmt.Println("")
				break
			default:
				break
			}
		}
	}

	//fmt.Println("Parking size is:", listParkedCars)
}

func getTotalCharges(durr int) int {
	if durr <= 2 {
		return 10
	} else {
		return 10 + (durr-2)*10
	}
}

func printStatus(listParkedCars map[string]*Car) {
	fmt.Println("Slot No. Registration No")

	keys := make([]string, 0, len(listParkedCars))
	for k := range listParkedCars {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, listParkedCars[k].Number)
	}
}
