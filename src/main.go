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
		if userInput == 7 {
			break
		}

		// Check if the option is valid.
		isValidMainOption = slices.Contains(validOptions, userInput)

		if isValidMainOption {
			// helpers.Clear()
			hasError = false

			if userInput == 1 {
				services.BinaryToDecimal()
			}

			if userInput == 2 {
				services.BinaryToOctal()
			}

			if userInput == 3 {
				services.BinaryToHexadecimal()
			}

			if userInput == 4 {
				services.DecimalToBinary()
			}

			if userInput == 5 {
				services.DecimalToOctal()
			}

			if userInput == 6 {
				services.DecimalToHexadecimal()
			}
		} else {
			hasError = true
		}
	}

	// Clear the terminal upon exit.
	helpers.Clear()
}
