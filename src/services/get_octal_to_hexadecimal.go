package services

import (
	"fmt"
	"strconv"
)

func getOctalToHexadecimal(octal int64) int64 {
	binaryStringValue := getOctalToBinary(octal)
	parsedBinary, err := strconv.ParseInt(binaryStringValue, 10, 64)

	if err != nil {
		return -1
	}

	fmt.Println(parsedBinary)

	return 86
}
