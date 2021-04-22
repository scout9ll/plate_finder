package search

import (
	"strconv"
	"strings"
)

func GetPlateListFromRangeString(rangeString string) ([]int, error) {
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
