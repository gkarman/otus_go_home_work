package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(textForUnpack string) (string, error) {
	var resultStringBuilder strings.Builder
	var prevRune rune
	length := len(textForUnpack)

	for i, currentRune := range textForUnpack {
		prevRuneIsDigit := unicode.IsDigit(prevRune)
		currentRuneIsDigit := unicode.IsDigit(currentRune)

		if prevRuneIsDigit && currentRuneIsDigit {
			return "", ErrInvalidString
		}

		if !prevRuneIsDigit && !currentRuneIsDigit {
			resultStringBuilder.WriteRune(prevRune)
			prevRune = currentRune

			isLastRune := i == length-1
			if isLastRune {
				resultStringBuilder.WriteRune(currentRune)
			}
			continue
		}

		if !prevRuneIsDigit {
			repeatCount, _ := strconv.Atoi(string(currentRune))
			stringForAdd := strings.Repeat(string(prevRune), repeatCount)
			resultStringBuilder.WriteString(stringForAdd)
			prevRune = currentRune
			continue
		}

		prevRune = currentRune

		isLastRune := i == length-1
		if isLastRune {
			resultStringBuilder.WriteRune(currentRune)
		}
	}

	textUnpack := resultStringBuilder.String()
	return textUnpack, nil
}
