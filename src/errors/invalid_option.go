package errors

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/helpers"
)

func InvalidOption() {
	fmt.Println(helpers.OneSpace, helpers.ColorRed, "Invalid option. Please try again.", helpers.ColorReset)
}
