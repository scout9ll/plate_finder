package search

import (
	"strconv"
	"strings"
)

func PartMatcher(plateNumber int, partNumber int) bool {
	plateString := strconv.Itoa(plateNumber)
	partString := strconv.Itoa(partNumber)
	return strings.Contains(plateString, partString)
}
