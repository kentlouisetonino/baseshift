package main

import (
	"fmt"
	"slices"

	"github.com/kentlouisetonino/baseshift/src/libs/computations"
	"github.com/kentlouisetonino/baseshift/src/libs/displays"
	"github.com/kentlouisetonino/baseshift/src/libs/errors"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	"github.com/kentlouisetonino/baseshift/src/libs/inputs"
)

func main() {
  // Variable declations with initialization.
  validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
  isValidMainOption := false
  hasError := false
  optionInput := 0
  
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
    userInput := inputs.MainOptionInput()
    isValidMainOption = slices.Contains(validOptions, userInput)
    
    if isValidMainOption {
      helpers.Clear()
      hasError = false;
      
      if (userInput == 1) {
        // This will handle the loop if option1 input is invalid.
        option1HasError := false
        
        for {
          // Clear first the terminal.
          helpers.Clear()

          // Display option 1 description.
          displays.Option1Description()
          helpers.AddNewLine()

          // Display error message.
          if (option1HasError) {
            errors.InvalidOption();
            helpers.AddNewLine()
          }

          // Display option 1 instruction.
          displays.Option1Instruction()
          helpers.AddNewLine()

          // Ask binary input.
          optionInput = inputs.Option1Input()
            
          if (optionInput == -1) {
            // Signal the app that there is an error.
            option1HasError = true

            continue
          } else {
            fmt.Println(helpers.TwoSpace, "Decimal :", computations.BinaryToDecimal(optionInput))
            
            // Reset the state
            option1HasError = false
            optionInput = 0

            break
          }

        }
      }

      break
    } else {
      hasError = true;
    }
  }
}

