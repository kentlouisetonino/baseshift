package services

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/kentlouisetonino/baseshift/src/displays"
	"github.com/kentlouisetonino/baseshift/src/errors"
	"github.com/kentlouisetonino/baseshift/src/helpers"
)

// Reversed the binary and return an array of binary.
func getReversedBinaryArray(binary int64) []int64 {
	var arrayOfBinary []int64

	for binary != 0 {
		arrayOfBinary = append(arrayOfBinary, binary%10)
		binary /= 10
	}

	return arrayOfBinary
}

// Reversed the array of int64.
func getReverseIntArray(array []int64) []int64 {
	lastIndex := len(array) - 1
	reversedArray := []int64{}

	for i := lastIndex; i >= 0; i-- {
		reversedArray = append(reversedArray, array[i])
	}

	return reversedArray
}

// This handles the binary input value and its validation.
func getBinaryInput() int64 {

	// Declare the variables.
	invalidNumbers := []int64{2, 3, 4, 5, 6, 7, 8, 9}
	var option int64
	var optionChecker string

	// Ask the option.
	fmt.Print(helpers.ThreeSpace, "Binary  : ")

	// Check if the string is a valid number.
	fmt.Scan(&optionChecker)
	parseOption, err := strconv.ParseInt(optionChecker, 10, 64)

	// If invalid numbers.
	if err != nil {
		return -1
	}

	// Assign the input the option variable.
	option = parseOption

	// Check if the input is a valid binary digit.
	isValid := true
	numberArray := getReversedBinaryArray(option)

	for i := 0; i < len(numberArray); i++ {
		isInvalid := slices.Contains(invalidNumbers, numberArray[i])

		if isInvalid {
			isValid = false
			break
		}

		continue
	}

	if isValid == false {
		return -1
	}

	return option
}

// Converts binary number to decimal.
func getDecimalValue(binary int64) int64 {
	binaryToArray := getReversedBinaryArray(binary)
	decimal := int64(0)

	for i := 0; i < len(binaryToArray); i++ {
		element := binaryToArray[i]
		decimal = decimal + (element * int64(math.Pow(2, float64(i))))
	}

	return decimal
}

// Converts binary number to octal.
func getOctalValue(binary int64) string {
	reversedBinaryArray := getReversedBinaryArray(binary)
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
	reversedArrayOfOctal := getReverseIntArray(arrayOfOctal)

	// Loop through the array, and concatenate it to the string.
	for i := 0; i < len(reversedArrayOfOctal); i++ {
		octal += strconv.Itoa(int(reversedArrayOfOctal[i]))
	}

	return octal
}

func BinaryToDecimal() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option1Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getBinaryInput()

		if optionInput == -1 {
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			// Display the output of computation for option 1.
			fmt.Println(helpers.TwoSpace, "Decimal :", getDecimalValue(optionInput))
			helpers.AddNewLine()

			// Ask user if want to try again, go back, or quit.
			// Any other key except 1 and 2 will be treated as quit.
			fmt.Print(helpers.TwoSpace, helpers.ColorGreen, " [1-Retry, 2-Back] : ", helpers.ColorReset)
			fmt.Scan(&optionMenu)

			if optionMenu == "1" {
				optionHasError = false
				continue
			} else if optionMenu == "2" {
				backToMainMenu = true
			} else {
				helpers.Exit()
			}
		}
	}
}

func BinaryToOctal() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 2 description.
		displays.Option2Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getBinaryInput()

		if optionInput == -1 {
			fmt.Println(optionInput)
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			// Display the output of computation for option 2.
			fmt.Println(helpers.TwoSpace, "Octal", helpers.OneSpace, ":", getOctalValue(optionInput))
			helpers.AddNewLine()

			// Ask user if want to try again, go back, or quit.
			// Any other key except 1 and 2 will be treated as quit.
			fmt.Print(helpers.TwoSpace, helpers.ColorGreen, " [1-Retry, 2-Back] : ", helpers.ColorReset)
			fmt.Scan(&optionMenu)

			if optionMenu == "1" {
				optionHasError = false
				continue
			} else if optionMenu == "2" {
				backToMainMenu = true
			} else {
				helpers.Exit()
			}
		}
	}
}

func BinaryToHexadecimal() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 2 description.
		displays.Option3Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getBinaryInput()

		if optionInput == -1 {
			fmt.Println(optionInput)
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			// Display the output of computation for option 2.
			fmt.Println(helpers.TwoSpace, "Hexadecimal", helpers.OneSpace, ":", getOctalValue(optionInput))
			helpers.AddNewLine()

			// Ask user if want to try again, go back, or quit.
			// Any other key except 1 and 2 will be treated as quit.
			fmt.Print(helpers.TwoSpace, helpers.ColorGreen, " [1-Retry, 2-Back] : ", helpers.ColorReset)
			fmt.Scan(&optionMenu)

			if optionMenu == "1" {
				optionHasError = false
				continue
			} else if optionMenu == "2" {
				backToMainMenu = true
			} else {
				helpers.Exit()
			}
		}
	}
}
