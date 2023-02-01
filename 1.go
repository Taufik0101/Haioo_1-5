package main

import (
	"fmt"
	"strconv"
)

type LastValue struct {
	Nominal string
	Jumlah  string
}

func main() {
	fmt.Print("Enter text: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	textNumber, _ := strconv.Atoi(input)

	CountNumber(textNumber)
}

func CountNumber(nominal int) {
	arrayMoney := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	var initial []LastValue
	indexMoney := 0
	initialMoney := nominal
	totalNominal := 0

	for next := true; next; next = indexMoney < len(arrayMoney) {
		if initialMoney >= 1 && initialMoney <= 99 {
			initialMoney = 100
		}

		if initialMoney >= arrayMoney[indexMoney] {
			initialMoney -= arrayMoney[indexMoney]
			totalNominal++
		} else {
			if totalNominal != 0 {
				initial = append(initial, LastValue{
					Nominal: strconv.Itoa(arrayMoney[indexMoney]),
					Jumlah:  strconv.Itoa(totalNominal),
				})
			}
			totalNominal = 0
			indexMoney++
		}
	}
	fmt.Println(initial)
}
