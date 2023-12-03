package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
    p "parser"
)

type Game = struct {
	id    int
	red   int
	green int
	blue  int
    power int64
}

func parseSet(line string, pos int) (int, int, int, int, error) {

    var red, green, blue int
    char, newPos, err := p.Peek(line, pos)
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
        num, newPos, err = p.ExpectNumber(line, newPos)
        if err != nil {
            return 0, 0, 0, newPos, fmt.Errorf("Invalid set: \n\t%s", err.Error())
        }
        color, newPos, err = p.ExpectOneOf(line, newPos, []string{" red"," blue", " green"})
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
        char, newPos, err = p.Peek(line, newPos)
        if err != nil || char == ';' || char == '\n' {
            newPos ++;
            break
        }
        _, newPos, err = p.ExpectWord(line, newPos, ", ")
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
        power: 0,
	}

    log.Printf("INFO: parsing line %s", line)
	_, pos, err := p.ExpectWord(line, 0, "Game ")
	if err != nil {
		return game, errors.New("Invalid line: Not a game data")
	}

	id, pos, err := p.ExpectNumber(line, pos)
	if err != nil {
		return game, fmt.Errorf("Invalid line: No game number found\n\t%s", err.Error())
	}
	game.id = id


    _, pos, err = p.ExpectWord(line, pos, ":")
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
	sum := 0
	powerSum := int64(0)
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
