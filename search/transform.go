package search

import (
	"strconv"
	"strings"
)

func plateFilter(rangeString string) bool {
	return strings.Contains(rangeString, "AD")
}

func GetPlateListFromRangeString(rangeString string) ([]int, error) {

	if !plateFilter(rangeString) {
		return []int{}, nil
	}

	fullPlaterRange := strings.Split(strings.ReplaceAll(rangeString, "AD", ""), " ~ ")
	var plateList []int
	startPlate, err := strconv.Atoi(fullPlaterRange[0])
	endPlate, err := strconv.Atoi(fullPlaterRange[1])
	if err != nil {
		return []int{}, err
	}
	for i := startPlate; i <= endPlate; i++ {
		plateList = append(plateList, i)
	}
	return plateList, nil
}
