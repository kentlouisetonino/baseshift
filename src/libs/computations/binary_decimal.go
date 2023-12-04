package computations

import (
	"math"
)

// Converts binary number to decimal.
func BinaryToDecimal(binary int) int {
  binaryToArray := NumberArray(binary)
  decimal := 0

  for i := 0; i < len(binaryToArray); i++ {
    element := binaryToArray[i]
    decimal = decimal + (element * int(math.Pow(2, float64(i))))
  }

  return decimal
}

