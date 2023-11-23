package main

import (
	"github.com/kentlouisetonino/baseshift/src/libs/displays"
	"github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func main() {
  // Clear up the screen first.
  helpers.Clear()

  // Display the app description.
  displays.AppDescription();
  helpers.NewLine();

  // Display the options.
  displays.AppOptions();
}

