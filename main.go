package main

import (
	"D-plate-finder/search"
	"fmt"
)

func main() {
	hitPlateList := search.SearchByPart(911)
	fmt.Printf("%v\n", hitPlateList)
}
