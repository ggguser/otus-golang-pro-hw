package hw02unpackstring

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func unpackAndJoin(pss ...packedSymbol) string {
	var sb strings.Builder
	for _, ps := range pss {
		sb.WriteString(ps.unpack())
	}
	return sb.String()
}

func isDigit(s string) (int, bool) {
	if d, err := strconv.Atoi(s); err == nil {
		return d, true
	}
	return 0, false
}

type packedSymbol struct {
	symbol rune
	reps   int
}

func (p packedSymbol) unpack() string {
	return strings.Repeat(string(p.symbol), p.reps)
}

func (p packedSymbol) String() string {
	return fmt.Sprintf("Symbol: %s, Reps: %d", string(p.symbol), p.reps)
}

func Unpack(s string) (string, error) {
	var packedSymbols []packedSymbol
	runes := []rune(s)
	slices.Reverse(runes)
	var seen bool
	for i, r := range runes {
		digit, ok := isDigit(string(r))
		if ok {
			if i == len(runes)-1 {
				return "", ErrInvalidString
			}

			symbol := runes[i+1]

			if _, ok = isDigit(string(symbol)); ok {
				return "", ErrInvalidString
			}
			packedSymbols = append(packedSymbols, packedSymbol{symbol, digit})
			seen = true
		} else {
			if !seen {
				packedSymbols = append(packedSymbols, packedSymbol{r, 1})
			}
			seen = false
		}
	}
	slices.Reverse(packedSymbols)
	return unpackAndJoin(packedSymbols...), nil
}
