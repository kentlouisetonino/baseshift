package services

import (
	"strconv"
	"strings"
)

func getDecimalToOctal(decimal int64) int64 {
	octalStringArray := []string{}

	// Divide the dividend until it reaches zero.
	for i := decimal; i != 0; i = i / 8 {
		// Append the reminder.
		// If no reminder add 0 in the beginning of array.
		// If contains a reminder add 1 in the beginning of array.
		octalStringArray = append([]string{strconv.Itoa(int(i - 8*(i/8)))}, octalStringArray...)
	}

	// Combine all the array of string of number.
	joinedBinaryStringArray := strings.Join(octalStringArray, "")

	// Convert the combines string of number into a number.
	parsedBinary, err := strconv.ParseInt(joinedBinaryStringArray, 10, 64)

	// If invalid numbers.
	if err != nil {
		return -1
	}

	return parsedBinary
}
