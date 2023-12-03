package parser
import (
    "errors"
    "fmt"
)
func ExpectWord(text string, pos int, prefix string) (string, int, error) {

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

func ExpectOneOf(text string, pos int, options []string) (string, int, error) {

	for _, option := range options {
		if _, newPos, err := ExpectWord(text, pos, option); err == nil {
			return option, newPos, err
		}
	}
	return text, pos, nil
}

func ExpectCharBetween(text string, pos int, low byte, high byte) (byte, int, error) {

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

func ExpectNumber(text string, pos int) (int, int, error) {
	num := 0

	if pos >= len(text) {
		return 0, pos, errors.New("No text left")
	}

	char, newPos, err := ExpectCharBetween(text, pos, '0', '9')
	if err != nil {
		return 0, pos, err
	}
	num = num*10 + int(char-'0')

	for {
		char, newPos, err = ExpectCharBetween(text, newPos, '0', '9')
		if err != nil {
			break
		}
		num = num*10 + int(char-'0')
	}

	return num, newPos, nil
}


func Peek(line string, pos int) (byte, int, error) {
    if pos >= len(line) {
        return 0, pos, fmt.Errorf("Not enough text")
    }
    return line[pos], pos, nil

}
