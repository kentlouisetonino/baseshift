package displays

import (
	"fmt"
	"strconv"

	helper "github.com/kentlouisetonino/baseshift/src/libs/helpers"
)

func openBracket() string {
  return helper.ColorBlue + "[" + helper.ColorReset
}

func closeBracket() string {
  return helper.ColorBlue + "]" + helper.ColorReset
}

func optionNumber(number int) string {
  return helper.ColorGreen + strconv.Itoa(number) + helper.ColorReset
}

func AppOptions() {
  fmt.Println(openBracket(), optionNumber(1), closeBracket(), "Binary to Decimal");
  fmt.Println(openBracket(), optionNumber(2), closeBracket(), "Binary to Octal");
  fmt.Println(openBracket(), optionNumber(3), closeBracket(), "Binary to Hexadecimal");
  fmt.Println(openBracket(), optionNumber(4), closeBracket(), "Decimal to Binary");
  fmt.Println(openBracket(), optionNumber(5), closeBracket(), "Decimal to Octal");
  fmt.Println(openBracket(), optionNumber(6), closeBracket(), "Decimal to Hexadecimal");
  fmt.Println(openBracket(), optionNumber(7), closeBracket(), "Octal to Binary");
  fmt.Println(openBracket(), optionNumber(8), closeBracket(), "Octal to Decimal");
  fmt.Println(openBracket(), optionNumber(9), closeBracket(), "Octal to Hexadecimal");
  fmt.Println(openBracket(), optionNumber(10), closeBracket(), "Hexadecimal to Binary");
  fmt.Println(openBracket(), optionNumber(11), closeBracket(), "Hexadecimal to Decimal");
  fmt.Println(openBracket(), optionNumber(12), closeBracket(), "Hexadecimal to Octal");
}

