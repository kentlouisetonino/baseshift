package main

import (
	"slices"

	"github.com/kentlouisetonino/baseshift/src/displays"
	"github.com/kentlouisetonino/baseshift/src/errors"
	"github.com/kentlouisetonino/baseshift/src/helpers"
	"github.com/kentlouisetonino/baseshift/src/services"
)

func main() {
	// Variable declations with initialization.
	validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	isValidMainOption := false
	hasError := false

	for {
		// Clear up the screen first.
		helpers.Clear()
		helpers.AddNewLine()

		// Display the app description.
		displays.AppDescription()
		helpers.AddNewLine()

		// Display error message.
		if hasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Display the options.
		displays.AppOptions()
		helpers.AddNewLine()

		// Ask input.
		userInput := helpers.MainOptionInput()

		// Exit the application.
		if userInput == 4 {
			break
		}

		// Check if the option is valid.
		isValidMainOption = slices.Contains(validOptions, userInput)

		if isValidMainOption {
			helpers.Clear()
			hasError = false

			// Handle the conversion from binary to decimal.
			if userInput == 1 {
				services.BinaryToDecimal()
			}

			// Handle the conversion from binary to octal.
			if userInput == 2 {
				services.BinaryToOctal()
			}

			// Handle the conversion from binary to hexadecimal.
			if userInput == 3 {
				services.BinaryToHexadecimal()
			}
		} else {
			hasError = true
		}
	}
}
