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
  validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
  isValidMainOption := false
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
    userInput := inputs.MainOptionInput()
    
    // Exit the application.
    if (userInput == 13) {
      break
    }

    isValidMainOption = slices.Contains(validOptions, userInput)
    
    if isValidMainOption {
      helpers.Clear()
      hasError = false;
      
      if (userInput == 1) {
        // This will handle the loop if option1 input is invalid.
        option1Menu := 0
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
          if (option1HasError) {
            errors.InvalidOption();
            helpers.AddNewLine()
          }

          // Ask binary input.
          option1Input = inputs.Option1Input()
            
          if (option1Input == -1) {
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
            
            if (option1Menu == 1) {
              option1HasError = false
              continue
            } else if (option1Menu == 2) {
              backToMainMenu = true
            } else {
              break
            }
          }
        }
      }
    } else {
      hasError = true;
    }
  }
}

