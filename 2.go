package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 2; i++ {
		fmt.Print("Enter Text " + strconv.Itoa(i) + " : ")
		scanner.Scan()
		text := scanner.Text()
		arr = append(arr, text)

	}
	Check(arr[0], arr[1])
}

func Check(string1 string, string2 string) {
	charsOri := []rune(string1)
	charsOri2 := []rune(string2)
	chars := []rune(string1)
	chars2 := []rune(string2)
	totalSama := 0

	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars2); j++ {
			if string(chars2[j]) == string(chars[i]) {
				totalSama++
				chars2 = remove(chars2, string(chars[i]))
			}
		}
	}

	if len(charsOri) == len(charsOri2) {
		if len(chars)-totalSama > 1 {
			fmt.Println("FALSE")
		} else {
			fmt.Println("TRUE")
		}
	} else if len(charsOri) > len(charsOri2) {
		if len(charsOri)-totalSama > 1 {
			fmt.Println("FALSE")
		} else {
			fmt.Println("TRUE")
		}
	} else if len(charsOri) < len(charsOri2) {
		if len(charsOri2)-totalSama > 1 {
			fmt.Println("FALSE")
		} else {
			fmt.Println("TRUE")
		}
	}
}

func remove(s []rune, r string) []rune {
	for i, v := range s {
		if string(v) == r {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
