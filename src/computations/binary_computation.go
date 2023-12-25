package computations

import (
	"math"
	"strconv"
)

// Converts binary number to decimal.
func BinaryToDecimal(binary int64) int64 {
	binaryToArray := ReversedBinaryArray(binary)
	decimal := int64(0)

	for i := 0; i < len(binaryToArray); i++ {
		element := binaryToArray[i]
		decimal = decimal + (element * int64(math.Pow(2, float64(i))))
	}

	return decimal
}

// Converts binary number to octal.
func BinaryToOctal(binary int64) string {
	reversedBinaryArray := ReversedBinaryArray(binary)
	octal := ""
	arrayOfOctal := []int64{}
	groupSum := int64(0)

	for i := 0; i < len(reversedBinaryArray); i++ {
		numWeight := i % 3
		isLastIndex := i == (len(reversedBinaryArray) - 1)

		// If remainder is 0, weight is 1.
		if numWeight%3 == 0 {
			groupSum += reversedBinaryArray[i] * 1

			// Push the value right away, if it is the last index.
			if isLastIndex {
				arrayOfOctal = append(arrayOfOctal, groupSum)
			}
		}

		// If remainder is 1, weight is 2.
		if numWeight%3 == 1 {
			groupSum += reversedBinaryArray[i] * 2

			// Push the value right away, if it is the last index.
			if isLastIndex {
				arrayOfOctal = append(arrayOfOctal, groupSum)
			}
		}

		// If remainder is 2, weight is 4.
		if numWeight%3 == 2 {
			arrayOfOctal = append(arrayOfOctal, groupSum+(reversedBinaryArray[i]*4))
			groupSum = 0
		}
	}

	// Reversed the octal array.
	reversedArrayOfOctal := ReverseIntArray(arrayOfOctal)

	// Loop through the array, and concatenate it to the string.
	for i := 0; i < len(reversedArrayOfOctal); i++ {
		octal += strconv.Itoa(int(reversedArrayOfOctal[i]))
	}

	return octal
}
