package hw02unpackstring

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(r rune) bool { return r >= '0' && r <= '9' }

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

func unpackAndJoin(pss ...packedSymbol) string {
	var sb strings.Builder
	for _, ps := range pss {
		sb.WriteString(ps.unpack())
	}
	return sb.String()
}

func Unpack(s string) (string, error) {
	var packedSymbols []packedSymbol
	runes := []rune(s)
	slices.Reverse(runes)
	var seen bool
	for i, r := range runes {
		if isDigit(r) {
			if i == len(runes)-1 {
				return "", ErrInvalidString
			}

			symbol := runes[i+1]

			if isDigit(symbol) {
				return "", ErrInvalidString
			}

			reps, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}

			packedSymbols = append(packedSymbols, packedSymbol{symbol, reps})
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
