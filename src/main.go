package main

import (
	"github.com/kentlouisetonino/baseshift/src/libs/display"
	"github.com/kentlouisetonino/baseshift/src/libs/helper"
)

func main() {
  // Clear up the screen first.
	helper.ClearScreen()

  // Display the app description.
  display.AppDescription();
}

