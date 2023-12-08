package inputs

import (
	"fmt"
	"slices"

	"github.com/kentlouisetonino/baseshift/src/libs/computations"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func Option1Input() int {
  
  // Declare the variables.
  invalidNumbers := []int{2, 3, 4, 5, 6, 7, 8 ,9}
  var option int;

  // Ask the option.
  fmt.Print(helpers.ThreeSpace, "Binary : ")

  // Assign the input the option variable.
  fmt.Scan(&option)
  
  // Check if the input is a valid binary digit.
  isValid := true
  numberArray := computations.NumberArray(option)

  for i := 0; i < len(numberArray); i++ {
    isInvalid := slices.Contains(invalidNumbers, numberArray[i])

    if (isInvalid) {
      isValid = false
      break
    }

    continue
  }

  if (isValid == false) {
    return -1
  }

  return option;
}
