package inputs

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func Option1Input() int {
  
  // Declare the variables.
  var option int;

  // Ask the option.
  fmt.Print(helpers.ThreeSpace, "Binary : ")

  // Assign the input the option variable.
  fmt.Scan(&option)

  return option;
}
