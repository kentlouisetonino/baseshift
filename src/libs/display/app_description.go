package display

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helper"
)

func AppDescription() {
  fmt.Println(helper.ColorBlue, "--------------------------------------", helper.ColorReset);
  helper.NewLine();
  fmt.Println(helper.ColorGreen, "       Welcome to BaseShift           ", helper.ColorReset);
  fmt.Println(helper.ColorWhite ,"                                      ", helper.ColorReset);
  fmt.Println(helper.ColorWhite, "  A CLI program that allows you to    ", helper.ColorReset);
  fmt.Println(helper.ColorWhite, "  convert a number system to another  ", helper.ColorReset);
  helper.NewLine();
  fmt.Println(helper.ColorBlue, "--------------------------------------", helper.ColorReset);
}

