package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/helpers"
)

func Option1Description() {
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorGreen, "       [1] Binary to Decimal          ", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
}
