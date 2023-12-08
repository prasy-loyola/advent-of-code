package parser

import (
	"bufio"
	"fmt"
)



type ParserError error
var EOL error = fmt.Errorf("Not enough text in line")
var EOF error = fmt.Errorf("Not enough text in file")
type Parser struct {
    Reader *bufio.Reader
    Pos int
    line string
}

func (p *Parser) ReadNextLine() (string, error) {
    if line, err := p.Reader.ReadBytes('\n'); err != nil {
        return "", EOF
    } else {
        p.line = string(line)
        p.Pos = 0
        return p.line, nil
    }
}

func (p *Parser) Skip(char byte) ([]byte, error) {
	if p.Pos >= len(p.line) {
		return nil, EOL
	}

    var bytes = make([]byte,0)
    for ; p.Pos < len(p.line); p.Pos++ {
        if p.line[p.Pos] != char {
            bytes = append(bytes, char)
            break
        }
    }
    return bytes, nil
}
func (p *Parser) ExpectWord(prefix string) (string, error) {

	if p.Pos >= len(p.line) || len(prefix) > len(p.line)+p.Pos {
        return "", fmt.Errorf("line: %s, pos:%d,  %w" ,p.line, p.Pos, EOL)
	}

	for i := 0; i < len(prefix); i++ {
		if p.line[i+p.Pos] != prefix[i] {
            return "", fmt.Errorf("Expected: '%s' but found character: '%c' at pos: '%d'", prefix, p.line[i+p.Pos], i+p.Pos)
		}
	}

    p.Pos += len(prefix)

	//log.Printf("INFO: found word '%s' at %d", prefix, pos)

	return prefix, nil
}

func (p *Parser) ExpectOneOf(options []string) (string, error) {

	for _, option := range options {
		if _, err := p.ExpectWord(option); err == nil {
			return option, err
		}
	}
	return "", nil
}

func (p *Parser) ExpectCharBetween(low byte, high byte) (byte, error) {

	if p.Pos >= len(p.line) {
		return 0, EOL
	}
	char := p.line[p.Pos]

	if char < low || char > high {
		return char, fmt.Errorf("Expected character between %c-%c but found '%c'",low, high, char)
	}
    p.Pos++
	//log.Printf("INFO: found char '%c' at %d", char, pos)
	return char, nil
}

func (p *Parser) ExpectNumber() (int, error) {
	num := 0

	if p.Pos >= len(p.line) {
		return 0, EOL
	}

	char, err := p.ExpectCharBetween('0', '9')
	if err != nil {
		return 0, err
	}
	num = num*10 + int(char-'0')

	for {
		char, err = p.ExpectCharBetween('0', '9')
		if err != nil {
			break
		}
		num = num*10 + int(char-'0')
	}

	return num, nil
}


func (p *Parser) Peek() (byte, error) {
    if p.Pos >= len(p.line) {
        return 0, EOL
    }
    return p.line[p.Pos], nil

}
