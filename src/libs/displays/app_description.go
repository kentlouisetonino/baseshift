package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	helper "github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func AppDescription() {
	fmt.Println(helpers.OneSpace, helper.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helper.AddNewLine()
	fmt.Println(helpers.OneSpace, helper.ColorGreen, "       Welcome to BaseShift           ", helpers.ColorReset)
	fmt.Println(helpers.OneSpace, helper.ColorWhite, "                                      ", helpers.ColorReset)
	fmt.Println(helpers.OneSpace, helper.ColorWhite, "  A CLI program that allows you to    ", helpers.ColorReset)
	fmt.Println(helpers.OneSpace, helper.ColorWhite, "  convert a number system to another  ", helpers.ColorReset)
	helper.AddNewLine()
	fmt.Println(helpers.OneSpace, helper.ColorBlue, "--------------------------------------", helpers.ColorReset)
}
