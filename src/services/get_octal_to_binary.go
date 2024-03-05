package services

func getOctalToBinary(octal int64) string {
	octalArray := getNumberArray(octal)
	binary := ""

	// Based the octal equivalent value in binary.
	octalSystem := map[int]string{
		0: "000",
		1: "001",
		2: "010",
		3: "011",
		4: "100",
		5: "101",
		6: "110",
		7: "111",
	}

	for i := 0; i < len(octalArray); i++ {
		binary = binary + octalSystem[int(octalArray[i])]
	}

	return binary
}
