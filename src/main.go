package main

import (
	"fmt"
	"slices"

	"github.com/kentlouisetonino/baseshift/src/libs/displays"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	"github.com/kentlouisetonino/baseshift/src/libs/inputs"
)

func main() {
  validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
  isValidOption := false
  // Clear up the screen first.
  helpers.Clear()

  // Display the app description.
  displays.AppDescription();
  helpers.AddNewLine();

  // Display the options.
  displays.AppOptions();
  helpers.AddNewLine();

  // Ask input.
  userInput := inputs.OptionInput()
  isValidOption = slices.Contains(validOptions, userInput)

  fmt.Println(validOptions)
  fmt.Println(userInput)
  fmt.Println(isValidOption)
}

