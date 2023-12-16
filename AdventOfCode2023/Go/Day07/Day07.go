package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hand struct {
	cards []int
	bet   int
}

func Parse(filename string, joker bool) []hand {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	hands := []hand{}

	y := 0
	for fscanner.Scan() {
		line := fscanner.Text()

		hand := hand{}
		hand.cards = []int{}

		for _, card := range string(strings.Split(line, " ")[0]) {
			cardStr := string(card)
			if cardStr == "T" {
				cardStr = "10"
			} else if cardStr == "J" {
				if joker {
					cardStr = "1"
				} else {
					cardStr = "11"
				}
			} else if cardStr == "Q" {
				cardStr = "12"
			} else if cardStr == "K" {
				cardStr = "13"
			} else if cardStr == "A" {
				cardStr = "14"
			}
			card, _ := strconv.Atoi(string(cardStr))
			hand.cards = append(hand.cards, card)
		}

		hand.bet, _ = strconv.Atoi(string(strings.Split(line, " ")[1]))
		hands = append(hands, hand)
		y++
	}

	return hands
}

func dup_count(list []int) map[int]int {

	duplicate_frequency := make(map[int]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func hand1IsBetter(hand1 hand, hand2 hand) bool {

	for i := 0; i < len(hand1.cards); i++ {
		if hand1.cards[i] > hand2.cards[i] {
			return true
		}
		if hand1.cards[i] < hand2.cards[i] {
			return false
		}
	}
	return false
}

func BubbleSort(handsGroups [][]hand) [][]hand {

	for a, _ := range handsGroups {
		array := handsGroups[a]
		for i := 0; i < len(array)-1; i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if hand1IsBetter(array[j], array[j+1]) {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
	}
	return handsGroups
}

func Part1(filename string) int {
	hands := Parse(filename, false)

	fiveOfAKind := []hand{}
	fourOfAKind := []hand{}
	fullHouse := []hand{}
	threeOfAKind := []hand{}
	twoPairs := []hand{}
	onePair := []hand{}
	highCard := []hand{}

	for _, hand := range hands {
		dup := dup_count(hand.cards)
		if len(dup) == 5 {
			highCard = append(highCard, hand)
		} else if len(dup) == 4 {
			onePair = append(onePair, hand)
		} else if len(dup) == 3 {
			for _, v := range dup {
				if v == 2 {
					twoPairs = append(twoPairs, hand)
					break
				} else if v == 3 {
					threeOfAKind = append(threeOfAKind, hand)
					break
				}
			}
		} else if len(dup) == 2 {
			for _, v := range dup {
				if v == 4 {
					fourOfAKind = append(fourOfAKind, hand)
					break
				} else if v == 3 {
					fullHouse = append(fullHouse, hand)
					break
				}
			}
		} else if len(dup) == 1 {
			fiveOfAKind = append(fiveOfAKind, hand)
		}

	}

	handsGroups := [][]hand{highCard, onePair, twoPairs, threeOfAKind, fullHouse, fourOfAKind, fiveOfAKind}
	BubbleSort(handsGroups)

	r := 1
	winnings := 0
	for g, _ := range handsGroups {
		for i, _ := range handsGroups[g] {
			winnings += handsGroups[g][i].bet * r
			r++
		}
	}
	return winnings
}

func Part2(filename string) int {
	hands := Parse(filename, true)

	fiveOfAKind := []hand{}
	fourOfAKind := []hand{}
	fullHouse := []hand{}
	threeOfAKind := []hand{}
	twoPairs := []hand{}
	onePair := []hand{}
	highCard := []hand{}

	for _, hand := range hands {
		dup := dup_count(hand.cards)

		numOfJokers := 0
		for c, v := range dup {
			if c == 1 {
				numOfJokers = v
			}
		}
		if numOfJokers == 0 {
			if len(dup) == 5 {
				highCard = append(highCard, hand)
			} else if len(dup) == 4 {
				onePair = append(onePair, hand)
			} else if len(dup) == 3 {
				for _, v := range dup {
					if v == 2 {
						twoPairs = append(twoPairs, hand)
						break
					} else if v == 3 {
						threeOfAKind = append(threeOfAKind, hand)
						break
					}
				}
			} else if len(dup) == 2 {
				for _, v := range dup {
					if v == 4 {
						fourOfAKind = append(fourOfAKind, hand)
						break
					} else if v == 3 {
						fullHouse = append(fullHouse, hand)
						break
					}
				}
			} else if len(dup) == 1 {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
		}

		if numOfJokers > 0 {
			if len(dup) == 5 {
				onePair = append(onePair, hand)
			} else if len(dup) == 4 {
				threeOfAKind = append(threeOfAKind, hand)
			} else if len(dup) == 3 && numOfJokers == 1 { 
				set := false
				for _, v := range dup {
					if v == 3 {
						fourOfAKind = append(fourOfAKind, hand)
						set = true
						break
					}
				}
				if !set {
					fullHouse = append(fullHouse, hand)
				}
			} else if len(dup) == 3 && numOfJokers > 1 {
				fourOfAKind = append(fourOfAKind, hand)
			} else if len(dup) <= 2 {
				fiveOfAKind = append(fiveOfAKind, hand)
			}
		}
	}

	handsGroups := [][]hand{highCard, onePair, twoPairs, threeOfAKind, fullHouse, fourOfAKind, fiveOfAKind}
	BubbleSort(handsGroups)

	r := 1
	winnings := 0
	for g, _ := range handsGroups {
		for i, _ := range handsGroups[g] {
			winnings += handsGroups[g][i].bet * r
			
			r++
		}
	}
	return winnings
}

func main() {
	fmt.Println("Advent-Of-Code 2023 - Day07")
	fmt.Printf("Tests : Answer to Part 1 = %v\n", Part1("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 1 = %v\n", Part1("inputs.txt"))
	fmt.Printf("Tests : Answer to Part 2 = %v\n", Part2("tests.txt"))
	fmt.Printf("Inputs: Answer to Part 2 = %v\n", Part2("inputs.txt"))
}
