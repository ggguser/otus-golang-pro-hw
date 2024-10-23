package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(s string) (int, bool) {
	if d, err := strconv.Atoi(s); err == nil {
		return d, true
	}
	return 0, false
}

func Unpack(s string) (string, error) {
	var unpacked string
	runes := []rune(s)
	for i, v := range runes {
		digit, ok := isDigit(string(v))
		if ok {
			if i == 0 {
				return "", ErrInvalidString
			}
			symbol := string(runes[i-1])

			if _, ok = isDigit(symbol); ok {
				return "", ErrInvalidString
			}
			repeated := strings.Repeat(symbol, digit)
			unpacked += repeated
		} else {
			unpacked += string(v)
		}
	}

	//utf8.RuneCountInString(repeated)
	//repeated := strings.Repeat(s, 3)

	return unpacked, nil
}
