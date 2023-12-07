package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand     string
	strength int
	bid      int
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIRS       = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

func main() {
	result1 := solution("input.txt", "23456789TJQKA", handStrength)
	println(result1)

	result2 := solution("input.txt", "J23456789TQKA", handStrength2)
	println(result2)
}

func solution(fileName string, cards string, strength_func func(hand string) int) int {
	var result = 0

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input = []Hand{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		hand, bid := line[0], line[1]
		b, _ := strconv.Atoi(bid)
		h := Hand{hand: hand, bid: b, strength: strength_func(hand)}
		input = append(input, h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	card_comparator := makeCardComparator(cards)
	sorter := handSorter(&input, card_comparator)
	sort.Slice(input, sorter)

	for i, h := range input {
		result += h.bid * (i + 1)
	}

	return result
}

func handSorter(input *[]Hand, comparator func(i, j byte) bool) func(i, j int) bool {
	return func(i, j int) bool {
		a, b := (*input)[i], (*input)[j]

		if a.strength == b.strength {
			for k := 0; k < len(a.hand); k += 1 {
				if a.hand[k] == b.hand[k] {
					continue
				}
				return comparator(a.hand[k], b.hand[k])
			}
		}
		return a.strength < b.strength
	}
}

func makeCardComparator(card_string string) func(i, j byte) bool {
	return func(i, j byte) bool {
		var cards = map[byte]int{}
		for i, v := range card_string {
			cards[byte(v)] = i
		}
		return cards[i] < cards[j]
	}
}

func handStrength2(hand string) int {

	counter := map[rune]int{}
	for _, c := range hand {
		counter[c]++
	}
	v := make([]int, 0, len(counter))
	for _, value := range counter {
		v = append(v, value)
	}
	sort.IntSlice(v).Sort()

	// Five of a kind AAAAA
	if len(v) == 1 {
		return FIVE_OF_A_KIND
	}

	// Four of a kind or full house AAAAB or AAABB
	if len(v) == 2 {

		// if there is joker in such hand its 5 of a kind AAAAJ or AAAJJ or AAJJJ
		if counter['J'] > 0 {
			return FIVE_OF_A_KIND
		}
		//Four of a kind AAAAB or AAAAB
		if v[0] == 1 {
			return FOUR_OF_A_KIND
		}
		//Full house AABBB or AAABB
		return FULL_HOUSE
	}

	// three of a kind or two pairs AAABC or AABBC
	if len(v) == 3 {

		// Three of a kind
		if v[len(v)-1] == 3 {
			// three of a kind can be 3 jokers, or 3 cards + 1 joker. In both cases it is four of a kind
			// AAABJ or ABJJJ
			if counter['J'] > 0 {
				return FOUR_OF_A_KIND
			}
			// AAABC
			return THREE_OF_A_KIND
		}

		//Two pairs can be 2 jokers + pair + random card, or two pairs + joker
		// AABBJ or AABJJ or AABBC
		// with 2 jokers it is four of a kind AABJJ
		if counter['J'] == 2 {
			return FOUR_OF_A_KIND
		}
		// two pairs + joker become full house AABBJ
		if counter['J'] == 1 {
			return FULL_HOUSE
		}
		// AABBC
		return TWO_PAIRS
	}

	// One pair AABCD
	if len(counter) == 4 {
		// one pair, can be 1 pair + 1 joker + 2 different cards
		// or pair of jokers + 3 different cards

		//if its pair of jokers hand becomes 3 of a kind ABCJJ
		if counter['J'] == 2 {
			return THREE_OF_A_KIND
		}
		// if it is a different pair then we have 1 joker and it is three of a kind
		// AABCJ
		if counter['J'] == 1 {
			return THREE_OF_A_KIND
		}
		return ONE_PAIR
	}

	// High card
	// with Joker it becomes a pair
	if counter['J'] > 0 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func handStrength(hand string) int {

	counter := map[rune]int{}

	for _, c := range hand {
		counter[c]++
	}

	v := make([]int, 0, len(counter))

	for _, value := range counter {
		v = append(v, value)
	}

	sort.IntSlice(v).Sort()

	if len(v) == 1 {
		return FIVE_OF_A_KIND
	}

	if len(v) == 4 {
		return ONE_PAIR
	}

	if len(v) == 3 {
		if v[len(v)-1] == 3 {
			return THREE_OF_A_KIND
		} else {
			return TWO_PAIRS
		}
	}

	if len(v) == 2 {
		if v[len(v)-1] == 4 {
			return FOUR_OF_A_KIND
		} else {
			return FULL_HOUSE
		}
	}
	return HIGH_CARD
}
