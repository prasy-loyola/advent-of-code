package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	parsermod "parser"
)

type HandType int

const (
	FiveOfaKind HandType = iota
	FourOfaKind
	FullHouse
	ThreeOfaKind
	TwoPair
	OnePair
	HighCard
)

//type HandType string
//const (
//    FiveOfaKind HandType = "FiveOfaKind"
//    FourOfaKind HandType = "FourOfaKind"
//    FullHouse HandType = "FullHouse"
//    ThreeOfaKind HandType = "ThreeOfaKind"
//    TwoPair HandType = "TwoPair"
//    OnePair HandType = "OnePair"
//    HighCard HandType = "HighCard"
//)

type Hand struct {
	cards string
	bid   int
	typ   HandType
	typ2  HandType
}

func getHandTypeFromFreqMap(freqmap [6]int) HandType {

	if freqmap[5] > 0 {
		return FiveOfaKind
	}
	if freqmap[4] > 0 {
		return FourOfaKind
	}
	if freqmap[3] > 0 {
		if freqmap[2] > 0 {
			return FullHouse
		}
		return ThreeOfaKind
	}
	if freqmap[2] > 0 {
		if freqmap[2] > 1 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard

}
func getHandType2(hand string) HandType {

	freq := make(map[byte]int)

	for i := 0; i < len(hand); i++ {
		if val, ok := freq[hand[i]]; !ok {
			freq[hand[i]] = 1
		} else {
			freq[hand[i]] = val + 1
		}
	}
	freqmap := [6]int{}

    if jokers, ok := freq['J']; ok {
	    maxFreq := 0
        maxFreqChar := byte('J')
	    for c, v := range freq {
	    	if c != 'J' {
                if v > maxFreq {
                    maxFreq = v
                    maxFreqChar = c
                }
	    	}
	    }
        if maxFreqChar != 'J' {
            freq[maxFreqChar] += jokers
            freq['J'] = 0
        }
    }

	for _, v := range freq {
		freqmap[v] += 1
	}

	return getHandTypeFromFreqMap(freqmap)

}

func getHandType(hand string) HandType {

	freq := make(map[byte]int)

	for i := 0; i < len(hand); i++ {
		if val, ok := freq[hand[i]]; !ok {
			freq[hand[i]] = 1
		} else {
			freq[hand[i]] = val + 1
		}
	}

	freqmap := [6]int{}
	for _, v := range freq {
		freqmap[v] += 1
	}

	return getHandTypeFromFreqMap(freqmap)
}

func less2(hands []Hand) func(i, j int) bool {
	lessFunc := func(i, j int) bool {
		hi := hands[i]
		hj := hands[j]
		if hi.typ2 != hj.typ2 {
			return hi.typ2 > hj.typ2
		}
		valueOf := func(c byte) int {
			if c > '1' && c <= '9' {
				return int(c - '0')
			}
			switch c {
			case 'T':
				return 10
			case 'J':
				return 1
			case 'Q':
				return 12
			case 'K':
				return 13
			case 'A':
				return 14
			}
			return -1
		}
		for i = 0; i < 5; i++ {
			ival := valueOf(hi.cards[i])
			jval := valueOf(hj.cards[i])
			if ival == jval {
				continue
			}
			return ival < jval
		}
		return false
	}
	return lessFunc
}

func less1(hands []Hand) func(i, j int) bool {
	lessFunc := func(i, j int) bool {
		hi := hands[i]
		hj := hands[j]
		if hi.typ != hj.typ {
			return hi.typ > hj.typ
		}
		valueOf := func(c byte) int {
			if c > '1' && c <= '9' {
				return int(c - '0')
			}
			switch c {
			case 'T':
				return 10
			case 'J':
				return 11
			case 'Q':
				return 12
			case 'K':
				return 13
			case 'A':
				return 14
			}
			return -1
		}
		for i = 0; i < 5; i++ {
			ival := valueOf(hi.cards[i])
			jval := valueOf(hj.cards[i])
			if ival == jval {
				continue
			}
			return ival < jval
		}
		return false
	}
	return lessFunc
}
func main() {
	fmt.Println("Hello World!")
	filepath := "input.txt"
	var input *os.File
	var err error
	if input, err = os.Open(filepath); err != nil {
		panic("Couldn't read input file " + filepath)
	}
	parser := parsermod.Parser{
		Reader: bufio.NewReader(input),
	}

	var line string

	hands := make([]Hand, 0)
	for {
		if err == parsermod.EOF {
			break
		} else if err != nil && err != parsermod.EOL {
			log.Fatalf("ERROR: Invalid format %+v", err)
		}
		if line, err = parser.ReadNextLine(); err != nil || line == "" {
			continue
		}

		//log.Printf("INFO: parsing line %s", line)
		var cards string
		if cards, err = parser.ReadString(' '); err != nil {
			continue
		}
		parser.Skip(' ')
		var bid int
		if bid, err = parser.ExpectNumber(); err != nil {
			continue
		}
		hands = append(hands, Hand{
			cards: cards,
			bid:   bid,
			typ:   getHandType(cards),
			typ2:  getHandType2(cards),
		})
	}

	sort.Slice(hands, less1(hands))

	//log.Printf("INFO: Hands %v", hands)

	puzzle1Result := 0
	for i, hand := range hands {
		puzzle1Result += hand.bid * (i + 1)
	}

	sort.Slice(hands, less2(hands))

	//log.Printf("INFO: Hands %v", hands)

	puzzle2Result := 0
	for i, hand := range hands {
		puzzle2Result += hand.bid * (i + 1)
	}

	log.Printf("INFO: puzzle1Result %d", puzzle1Result)
	log.Printf("INFO: puzzle2Result %d", puzzle2Result)
}
