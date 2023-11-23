package displays

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	helper "github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func AppDescription() {
  fmt.Println(helper.ColorBlue, "--------------------------------------", helpers.ColorReset);
  helper.NewLine();
  fmt.Println(helper.ColorGreen, "       Welcome to BaseShift           ", helpers.ColorReset);
  fmt.Println(helper.ColorWhite ,"                                      ", helpers.ColorReset);
  fmt.Println(helper.ColorWhite, "  A CLI program that allows you to    ", helpers.ColorReset);
  fmt.Println(helper.ColorWhite, "  convert a number system to another  ", helpers.ColorReset);
  helper.NewLine();
  fmt.Println(helper.ColorBlue, "--------------------------------------", helpers.ColorReset);
}

