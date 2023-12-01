package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digits := map[rune]string{'0': "0", '1': "1", '2': "2", '3': "3", '4': "4", '5': "5", '6': "6", '7': "7", '8': "8", '9': "9"}
	digitsAsLetters := map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	for key, val := range digitsAsLetters {
		temp := ""
		for i := 0; i < len(key)-1; i++ {
			temp += string(key[i])
			digitsAsLetters[temp] = "k"
		}

		temp = ""
		for i := len(key) - 1; i >= 0; i-- {
			temp += string(key[i])
			if i == 0 {
				digitsAsLetters[temp] = val
			} else {
				digitsAsLetters[temp] = "k"
			}
		}
	}

	fmt.Println(digitsAsLetters)

	currVal := 0
	for scanner.Scan() {
		num := ""
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if val, ok := digits[rune(line[i])]; ok {
				num += val
				break
			} else {
				tempString := ""
				temp := i
				boo := true
				for boo {
					tempString += string(line[temp])
					temp += 1
					val, ok := digitsAsLetters[tempString]
					boo = ok
					if ok && val != "k" {
						num += val
						break
					}
				}
				if num != "" {
					break
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if val, ok := digits[rune(line[i])]; ok {
				num += val
				break
			} else {
				tempString := ""
				temp := i
				boo := true
				for boo {
					tempString += string(line[temp])
					temp -= 1
					val, ok := digitsAsLetters[tempString]
					boo = ok
					if ok && val != "k" {
						num += val
						break
					}
				}
				if len(num) != 1 {
					break
				}
			}
		}
		val, _ := strconv.Atoi(num)
		currVal += val

	}
	fmt.Println(currVal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
