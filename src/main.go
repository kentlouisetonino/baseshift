package main

import (
	"fmt"
	"slices"

	"github.com/kentlouisetonino/baseshift/src/libs/displays"
	"github.com/kentlouisetonino/baseshift/src/libs/errors"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	"github.com/kentlouisetonino/baseshift/src/libs/inputs"
)

func main() {
  // Variable declations with initialization.
  validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
  isValidOption := false
  hasError := false
  
  for {
    // Clear up the screen first.
    helpers.Clear()

    // Display the app description.
    displays.AppDescription();
    helpers.AddNewLine();
    
    // Display error message.
    if (hasError) {
      errors.InvalidOption();
      helpers.AddNewLine()
    }

    // Display the options.
    displays.AppOptions();
    helpers.AddNewLine();

    // Ask input.
    userInput := inputs.OptionInput()
    isValidOption = slices.Contains(validOptions, userInput)
    
    if isValidOption {
      helpers.AddNewLine()
      hasError = false;
      fmt.Println(helpers.TwoSpace, "Thank you.")
      break
    } else {
      hasError = true;
    }
  }
}

