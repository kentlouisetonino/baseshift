package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	helper "github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func Option1Description() {
	fmt.Println(helpers.OneSpace, helper.ColorBlue, "--------------------------------------", helpers.ColorReset)
	helper.AddNewLine()
	fmt.Println(helpers.OneSpace, helper.ColorGreen, "       [1] Binary to Decimal          ", helpers.ColorReset)
	helper.AddNewLine()
	fmt.Println(helpers.OneSpace, helper.ColorBlue, "--------------------------------------", helpers.ColorReset)
}
