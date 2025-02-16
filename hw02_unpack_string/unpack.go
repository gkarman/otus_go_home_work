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
		isLastRune := i == length-1

		if prevRune == 0 && currentRuneIsDigit {
			return "", ErrInvalidString
		}

		if prevRuneIsDigit && currentRuneIsDigit {
			return "", ErrInvalidString
		}

		if !prevRuneIsDigit && !currentRuneIsDigit {
			if prevRune != 0 {
				resultStringBuilder.WriteRune(prevRune)
			}
			prevRune = currentRune

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
		if isLastRune {
			resultStringBuilder.WriteRune(currentRune)
		}
	}

	textUnpack := resultStringBuilder.String()
	return textUnpack, nil
}
