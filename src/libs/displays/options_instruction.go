package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func Option1Instruction() {
  fmt.Println(helpers.OneSpace, helpers.ColorYellow, "Please input a valid binary numbers.", helpers.ColorReset)
}
