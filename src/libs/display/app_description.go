package display

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/helper"
)

func AppDescription() {
  fmt.Println("\033[34m", "--------------------------------------", "\033[0m");
  helper.NewLine();
  fmt.Println("\033[32m", "       Welcome to BaseShift           ", "\033[0m");
  fmt.Println("\033[97m" ,"                                      ", "\033[0m");
  fmt.Println("\033[97m", "  A CLI program that allows you to    ", "\033[0m");
  fmt.Println("\033[97m", "  convert a number system to another  ", "\033[0m");
  helper.NewLine();
  fmt.Println("\033[34m", "--------------------------------------", "\033[0m");
}

