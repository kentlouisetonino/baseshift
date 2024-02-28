package services

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/displays"
	"github.com/kentlouisetonino/baseshift/src/errors"
	"github.com/kentlouisetonino/baseshift/src/helpers"
)

func DecimalToBinary() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option4Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getDecimalInput()

		if optionInput == -1 {
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			fmt.Println(helpers.TwoSpace, "Binary", helpers.FiveSpace, ":", getDecimalToBinary(optionInput))
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
				helpers.Clear()
				helpers.Exit()
			}
		}
	}
}

func DecimalToOctal() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option5Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getDecimalInput()

		if optionInput == -1 {
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			fmt.Println(helpers.TwoSpace, "Octal", helpers.SixSpace, ":", getDecimalToOctal(optionInput))
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
				helpers.Clear()
				helpers.Exit()
			}
		}
	}
}

func DecimalToHexadecimal() {
	// This will handle the loop if option1 input is invalid.
	optionMenu := "0"
	optionInput := int64(0)
	optionHasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option6Description()
		helpers.AddNewLine()

		// Display error message.
		if optionHasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		optionInput = getDecimalInput()

		if optionInput == -1 {
			// Signal the app that there is an error.
			optionHasError = true

			continue
		} else {
			fmt.Println(helpers.TwoSpace, "Hexadecimal ", ":", getDecimalToHexadecimal(optionInput))
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
				helpers.Clear()
				helpers.Exit()
			}
		}
	}
}
