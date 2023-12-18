package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type hand struct{
	rank int
	cards string
	bet int
}


func main() {
	 
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rankCards := func(cards string) int {
		sortedHand := map[byte]int{}
		for i :=0 ; i < len(cards); i++{
			sortedHand[cards[i]] += 1 
		}
		if len(sortedHand) == 1{
			return 7 
		}
		if len(sortedHand) == 2 && (sortedHand[cards[0]] == 4  || sortedHand[cards[1]] == 4){
			return 6 
		} 
		for _,val:= range sortedHand{
			if val == 3 {
				if len(sortedHand) == 2{
					return 5
				}
				return 4
			}
		}

		for _,val:= range sortedHand{
			if val == 2{
				if len(sortedHand) == 3{
					return 3
				}
				return 2
			}
		}
		return 1

	}

	cardrank := map[rune]int{'A' : 14, 'K' : 13, 'Q' : 12, 'J' : 11, 'T' : 10, '9' : 9, '8' : 8, '7' : 7, '6' : 6, '5' : 5, '4' : 4, '3' : 3, '2' : 2}
	hands := []hand{}
	for i := 0; i < len(lines); i++{
		split := strings.Split(lines[i], " ")
		num, _ := strconv.Atoi(split[1])
		hands = append(hands, hand{rankCards(split[0]), split[0], num})
	}

	slices.SortFunc(hands, func(a, b hand) int {

		if a.rank > b.rank{
			return 1
		}
		if a.rank < b.rank{
			return -1 
		} 

		for i := 0; i < len(a.cards); i ++{

			if cardrank[rune(a.cards[i])]> cardrank[rune(b.cards[i])]{
				return 1
			}

			if cardrank[rune(a.cards[i])] < cardrank[rune(b.cards[i])]{
				return -1
			}
		}  
		return 0
	})

	tot := 0
	for i := 0; i < len(hands); i++{
		tot += (i + 1) * hands[i].bet
	}

	fmt.Println(tot)
	
}
