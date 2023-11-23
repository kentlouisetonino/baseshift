package main

import (
	"fmt"

	"github.com/kentlouisetonino/baseshift/src/libs/displays"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
	"github.com/kentlouisetonino/baseshift/src/libs/inputs"
)

func main() {
  // Clear up the screen first.
  helpers.Clear()

  // Display the app description.
  displays.AppDescription();
  helpers.AddNewLine();

  // Display the options.
  displays.AppOptions();
  helpers.AddNewLine();

  // Ask input.
  fmt.Println(inputs.OptionInput())
}

