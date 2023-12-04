package inputs

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func MainOptionInput() int {
  // Declare the variables.
  var option int;

  // Ask the option.
  fmt.Print(helpers.ThreeSpace, "Please choose an option: ")

  // Assign the input the option variable.
  fmt.Scan(&option)

  return option;
}

