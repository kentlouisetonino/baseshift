package services

import (
	"strconv"
)

func getOctalToHexadecimal(octal int64) string {
	// Get the equivalent of octal in binary.
	binaryStringValue := getOctalToBinary(octal)

	// Convert the string binary value to int64.
	parsedBinary, err := strconv.ParseInt(binaryStringValue, 10, 64)

	if err != nil {
		return ""
	}

	// Convert the int64 value to int64 array and then reverse.
	binaryArray := getNumberArray(parsedBinary)

	// Check if the length of the binary array is divisible by 4.
	isDivisibleByFour := (len(binaryArray) % 4) == 0

	groupHexadecimal := ""
	finalHexadecimal := ""

	if isDivisibleByFour {
		for i := 0; i < len(binaryArray)+1; i++ {
			if len(groupHexadecimal) == 4 {
				switch groupHexadecimal {
				case "0000":
					finalHexadecimal = finalHexadecimal + "0"
					groupHexadecimal = ""
				case "0001":
					finalHexadecimal = finalHexadecimal + "1"
					groupHexadecimal = ""
				case "0010":
					finalHexadecimal = finalHexadecimal + "2"
					groupHexadecimal = ""
				case "0011":
					finalHexadecimal = finalHexadecimal + "3"
					groupHexadecimal = ""
				case "0100":
					finalHexadecimal = finalHexadecimal + "4"
					groupHexadecimal = ""
				case "0101":
					finalHexadecimal = finalHexadecimal + "5"
					groupHexadecimal = ""
				case "0110":
					finalHexadecimal = finalHexadecimal + "6"
					groupHexadecimal = ""
				case "0111":
					finalHexadecimal = finalHexadecimal + "7"
					groupHexadecimal = ""
				case "1000":
					finalHexadecimal = finalHexadecimal + "8"
					groupHexadecimal = ""
				case "1001":
					finalHexadecimal = finalHexadecimal + "9"
					groupHexadecimal = ""
				case "1010":
					finalHexadecimal = finalHexadecimal + "A"
					groupHexadecimal = ""
				case "1011":
					finalHexadecimal = finalHexadecimal + "B"
					groupHexadecimal = ""
				case "1100":
					finalHexadecimal = finalHexadecimal + "C"
					groupHexadecimal = ""
				case "1101":
					finalHexadecimal = finalHexadecimal + "D"
					groupHexadecimal = ""
				case "1110":
					finalHexadecimal = finalHexadecimal + "E"
					groupHexadecimal = ""
				case "1111":
					finalHexadecimal = finalHexadecimal + "F"
					groupHexadecimal = ""
				}

			} else {
				if i != 8 {
					if binaryArray[i] == 0 {
						groupHexadecimal = groupHexadecimal + "0"
					} else {
						groupHexadecimal = groupHexadecimal + "1"
					}
				}
			}
		}
	}

	return finalHexadecimal
}
