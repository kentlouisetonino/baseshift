package services

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/displays"
	"github.com/kentlouisetonino/baseshift/src/helpers"
)

func DecimalToBinary() {
	backToMainMenu := false
	var test string

	for backToMainMenu == false {
		// Clear first the terminal.
		helpers.Clear()

		// Display option 2 description.
		displays.Option4Description()
		helpers.AddNewLine()

		fmt.Scan(&test)

		backToMainMenu = true
	}
}
