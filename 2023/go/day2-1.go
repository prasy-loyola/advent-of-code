package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type TokenType int8

const (
	Number TokenType = iota
	TokenList
	Word
)

type Token = struct {
	tType TokenType
	text  string
}

type Game = struct {
	id    int
	red   int
	green int
	blue  int
}

// expect("Game"), expect(" "), expect(Number), expect(": "), expect(Number), expect(("red", "green", "blue")), expect(";")

func expectWord(text string, pos int, prefix string) (string, int, error) {

	if pos >= len(text) || len(prefix) > len(text)+pos {
		return "", pos, errors.New("Not enough text")
	}

	for i := 0; i < len(prefix); i++ {
		if text[i+pos] != prefix[i] {
            return "", pos, fmt.Errorf("Expected: '%s' but found character: '%c' at pos: '%d'", prefix, text[i+pos], i+pos)
		}
	}

	//log.Printf("INFO: found word '%s' at %d", prefix, pos)

	return prefix, pos + len(prefix), nil
}

func expectOneOf(text string, pos int, options []string) (string, int, error) {

	for _, option := range options {
		if _, newPos, err := expectWord(text, pos, option); err == nil {
			return option, newPos, err
		}
	}
	return text, pos, nil
}

func expectCharBetween(text string, pos int, low byte, high byte) (byte, int, error) {

	if pos >= len(text) {
		return 0, pos, errors.New("No text left")
	}
	char := text[pos]

	if char < low || char > high {
		return char, pos, fmt.Errorf("Expected Number but found '%c'", char)
	}

	//log.Printf("INFO: found char '%c' at %d", char, pos)
	return char, pos + 1, nil
}

func expectNumber(text string, pos int) (int, int, error) {
	num := 0

	if pos >= len(text) {
		return 0, pos, errors.New("No text left")
	}

	char, newPos, err := expectCharBetween(text, pos, '0', '9')
	if err != nil {
		return 0, pos, err
	}
	num = num*10 + int(char-'0')

	for {
		char, newPos, err = expectCharBetween(text, newPos, '0', '9')
		if err != nil {
			break
		}
		num = num*10 + int(char-'0')
	}

	return num, newPos, nil
}


func peek(line string, pos int) (byte, int, error) {
    if pos >= len(line) {
        return 0, pos, fmt.Errorf("Not enough text")
    }
    return line[pos], pos, nil

}
func parseSet(line string, pos int) (int, int, int, int, error) {

    var red, green, blue int
    char, newPos, err := peek(line, pos)
    if err != nil {
        return 0, 0, 0, pos, fmt.Errorf("Couldn't parse Set\n\t%s", err.Error())
    }

    if char == ' ' {
        newPos++
    }


    var num int
    var color string

    for {
        if pos > len(line) {
            break
        }
        num, newPos, err = expectNumber(line, newPos)
        if err != nil {
            return 0, 0, 0, newPos, fmt.Errorf("Invalid set: \n\t%s", err.Error())
        }
        color, newPos, err = expectOneOf(line, newPos, []string{" red"," blue", " green"})
        switch color {
        case " red":
            red = num
            break
        case " blue":
            blue = num
            break
        case " green":
            green = num
            break
        }
        char, newPos, err = peek(line, newPos)
        if err != nil || char == ';' || char == '\n' {
            newPos ++;
            break
        }
        _, newPos, err = expectWord(line, newPos, ", ")
        if err != nil {
            return 0, 0, 0, newPos, fmt.Errorf("Invalid set: \n\t%s", err.Error())
        }
    }

    return red, green, blue, newPos, nil

}

func parseGame(line string) (Game, error) {
	game := Game{
		id:    0,
		red:   0,
		green: 0,
		blue:  0,
	}

    log.Printf("INFO: parsing line %s", line)
	_, pos, err := expectWord(line, 0, "Game ")
	if err != nil {
		return game, errors.New("Invalid line: Not a game data")
	}

	id, pos, err := expectNumber(line, pos)
	if err != nil {
		return game, fmt.Errorf("Invalid line: No game number found\n\t%s", err.Error())
	}
	game.id = id


    _, pos, err = expectWord(line, pos, ":")
	if err != nil {
		return game, fmt.Errorf("Invalid line: \n\t%s", err.Error())
	}

    var red, green, blue int

    
    var newPos = pos
    for {
        red, green, blue, newPos, err = parseSet(line, newPos)
        //log.Printf("red: %d, green: %d, blue: %d", red, green, blue)

        if err != nil {
            //log.Printf("DEBUG: error in parsing \n\t%s" , err.Error())
            break;
        }
        if red > game.red {
            game.red = red
        }
        if green > game.green {
            game.green = green
        }
        if blue > game.blue {
            game.blue = blue
        }
    }

	return game, nil
}

func main() {
	filepath := "day2-1.input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		panic("Not able to read file")
	}

	reader := bufio.NewReader(input)
	sum := 0
	expectedRed := 12
	expectedGreen := 13
	expectedBlue := 14

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		game, err := parseGame(string(line))

		if err != nil {
			fmt.Println(err)
			panic("Unable to continue")
		}
		fmt.Println(game)

		if game.red <= expectedRed &&
			game.green <= expectedGreen &&
			game.blue <= expectedBlue {
			sum += game.id
		}

	}

    log.Printf("INFO: Answer %d", sum)

}
