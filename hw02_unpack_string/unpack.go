package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(textForUnpack string) (string, error) {
	var resultStringBuilder strings.Builder
	var prevRune rune
	length := utf8.RuneCountInString(textForUnpack)
	currentIndex := 0
	for _, currentRune := range textForUnpack {
		if !isValidString(currentIndex, prevRune, currentRune) {
			return "", ErrInvalidString
		}

		repeatCount, err := getCountPreviousRepeats(prevRune, currentRune)
		if err != nil {
			return "", ErrInvalidString
		}

		stringForAdd := strings.Repeat(string(prevRune), repeatCount)
		resultStringBuilder.WriteString(stringForAdd)
		prevRune = currentRune

		isLastRune := currentIndex == length-1
		if isLastRune && !unicode.IsDigit(currentRune) {
			resultStringBuilder.WriteString(string(currentRune))
		}

		currentIndex++
	}

	return resultStringBuilder.String(), nil
}

func isValidString(index int, prevRune rune, currentRune rune) bool {
	prevRuneIsDigit := unicode.IsDigit(prevRune)
	currentRuneIsDigit := unicode.IsDigit(currentRune)
	if index == 0 && currentRuneIsDigit {
		return false
	}
	if prevRuneIsDigit && currentRuneIsDigit {
		return false
	}
	return true
}

func getCountPreviousRepeats(prevRune rune, currentRune rune) (int, error) {
	if prevRune == '\x00' {
		return 0, nil
	}

	if unicode.IsDigit(prevRune) {
		return 0, nil
	}

	if !unicode.IsDigit(currentRune) {
		return 1, nil
	}

	number, err := strconv.Atoi(string(currentRune))
	if err != nil {
		return 0, err
	}

	return number, nil
}
