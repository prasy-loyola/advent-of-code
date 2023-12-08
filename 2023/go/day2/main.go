package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
    parsermod "parser"
)

type Game = struct {
	id    int
	red   int
	green int
	blue  int
    power int64
}

func parseSet(p *parsermod.Parser) (int, int, int, error) {

    var red, green, blue int
    char, err := p.Peek()
    if err != nil {
        return 0, 0, 0, fmt.Errorf("Couldn't parse Set\n\t%w", err)
    }

    if char == ' ' {
        p.Pos++
    }


    var num int
    var color string

    for {
        _, err = p.ExpectWord("\n")
        if err != nil && err == parsermod.EOL { break }
        num, err = p.ExpectNumber()
        if err != nil {
            return 0, 0, 0, fmt.Errorf("Invalid set: \n\t%s", err.Error())
        }
        color, err = p.ExpectOneOf([]string{" red"," blue", " green"})
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
        char, err = p.Peek()
        if err != nil || char == ';' || char == '\n' {
            p.Pos++;
            break
        }
        _, err = p.ExpectWord(", ")
        if err != nil {
            return 0, 0, 0, fmt.Errorf("Invalid set: \n\t%s", err.Error())
        }
    }

    return red, green, blue, nil

}

func parseGame(p *parsermod.Parser) (Game, error) {
	game := Game{
		id:    0,
		red:   0,
		green: 0,
		blue:  0,
        power: 0,
	}

    //log.Printf("INFO: parsing line %s", line)
	_, err := p.ExpectWord("Game ")
	if err != nil {
		return game, errors.New("Invalid line: Not a game data")
	}

	id, err := p.ExpectNumber()
	if err != nil {
		return game, fmt.Errorf("Invalid line: No game number found\n\t%s", err.Error())
	}
	game.id = id


    _, err = p.ExpectWord(":")
	if err != nil {
		return game, fmt.Errorf("Invalid line: \n\t%s", err.Error())
	}

    var red, green, blue int

    
    for {
        red, green, blue, err = parseSet(p)
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
    game.power = int64(game.red) * int64(game.green) * int64(game.blue)

	return game, nil
}

func main() {
	filepath := "input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		panic("Not able to read file")
	}

	reader := bufio.NewReader(input)
    parser := parsermod.Parser {
        Reader: reader,
    }
	sum := 0
	powerSum := int64(0)
	expectedRed := 12
	expectedGreen := 13
	expectedBlue := 14

    var line string
	for {
        if err != nil && err != parsermod.EOL {
            log.Fatalf("ERROR: invalid line %s\n%s", line, err.Error())

        }
        if line, err = parser.ReadNextLine(); err != nil {
			break
		}

		game, err := parseGame(&parser)

		if err != nil {
			fmt.Println(err)
			panic("Unable to continue")
		}
		fmt.Println(game)


        powerSum += game.power

		if game.red <= expectedRed &&
			game.green <= expectedGreen &&
			game.blue <= expectedBlue {
			sum += game.id
		}

	}

    log.Printf("INFO: Answer %d", sum)
    log.Printf("INFO: Answer Total power %d", powerSum)

}
