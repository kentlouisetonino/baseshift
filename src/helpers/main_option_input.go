package helpers

import (
	"fmt"
)

func MainOptionInput() int {
	// Declare the variables.
	var option int

	// Ask the option.
	fmt.Print(ThreeSpace, "Please choose an option: ")

	// Assign the input the option variable.
	fmt.Scan(&option)

	return option
}
