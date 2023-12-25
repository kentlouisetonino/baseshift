package services

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/computations"
	"github.com/kentlouisetonino/baseshift/src/displays"
	"github.com/kentlouisetonino/baseshift/src/errors"
	"github.com/kentlouisetonino/baseshift/src/helpers"
	"github.com/kentlouisetonino/baseshift/src/inputs"
)

func BinaryToDecimal() {
	// This will handle the loop if option1 input is invalid.
	option1Menu := "0"
	option1Input := int64(0)
	option1HasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option1Description()
		helpers.AddNewLine()

		// Display error message.
		if option1HasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		option1Input = inputs.BinaryInput()

		if option1Input == -1 {
			fmt.Println(option1Input)
			// Signal the app that there is an error.
			option1HasError = true

			continue
		} else {
			// Display the output of computation ffor option 1.
			fmt.Println(helpers.TwoSpace, "Decimal :", computations.BinaryToDecimal(option1Input))
			helpers.AddNewLine()

			// Ask user if want to try again, go back, or quit.
			// Any other key except 1 and 2 will be treated as quit.
			fmt.Print(helpers.TwoSpace, helpers.ColorGreen, " [1-Retry, 2-Back] : ", helpers.ColorReset)
			fmt.Scan(&option1Menu)

			if option1Menu == "1" {
				option1HasError = false
				continue
			} else if option1Menu == "2" {
				backToMainMenu = true
			} else {
				helpers.Exit()
			}
		}
	}
}

func BinaryToOctal() {
	// This will handle the loop if option1 input is invalid.
	option1Menu := "0"
	option1Input := int64(0)
	option1HasError := false
	backToMainMenu := false

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 1 description.
		displays.Option2Description()
		helpers.AddNewLine()

		// Display error message.
		if option1HasError {
			errors.InvalidOption()
			helpers.AddNewLine()
		}

		// Ask binary input.
		option1Input = inputs.BinaryInput()

		if option1Input == -1 {
			fmt.Println(option1Input)
			// Signal the app that there is an error.
			option1HasError = true

			continue
		} else {
			// Display the output of computation ffor option 1.
			fmt.Println(helpers.TwoSpace, "Octal", helpers.OneSpace, ":", computations.BinaryToDecimal(option1Input))
			helpers.AddNewLine()

			// Ask user if want to try again, go back, or quit.
			// Any other key except 1 and 2 will be treated as quit.
			fmt.Print(helpers.TwoSpace, helpers.ColorGreen, " [1-Retry, 2-Back] : ", helpers.ColorReset)
			fmt.Scan(&option1Menu)

			if option1Menu == "1" {
				option1HasError = false
				continue
			} else if option1Menu == "2" {
				backToMainMenu = true
			} else {
				helpers.Exit()
			}
		}
	}
}
