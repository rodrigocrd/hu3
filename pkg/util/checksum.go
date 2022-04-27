package util

import "strconv"

func Checksum(rounds [][]int, baseDocument string) int {

	checksumValue := 0

	for _, weights := range rounds {
		checksumValue = (checksumValue * 10) + checksum(weights, baseDocument)
		baseDocument += strconv.Itoa(checksumValue)
	}

	return checksumValue
}

func checksum(weights []int, baseDocument string) int {
	checksumValue := 0

	for i, weight := range weights {
		singleNumber := string(baseDocument[i])
		intValue, err := strconv.Atoi(singleNumber)

		if err != nil {
			//TODO handle
			panic(err)
		}

		checksumValue += intValue * weight
	}

	mod := checksumValue % 11

	if mod < 2 {
		return 0
	}

	return 11 - mod
}
