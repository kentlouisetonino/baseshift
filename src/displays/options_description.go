package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/helpers"
)

func Option1Description() {
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorGreen, "       [1] Binary to Decimal         ", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
}

func Option2Description() {
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorGreen, "       [2] Binary to Octal           ", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
}

func Option3Description() {
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorGreen, "       [3] Binary to Hexadecimal     ", helpers.ColorReset)
	helpers.AddNewLine()
	fmt.Println(helpers.OneSpace, helpers.ColorBlue, "--------------------------------------", helpers.ColorReset)
}
