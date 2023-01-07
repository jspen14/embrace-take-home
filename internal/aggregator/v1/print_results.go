package v1_aggregator

import "fmt"

func (a *v1Aggregator) PrintResults() {

	for day, dayStreaks := range a.aggregates {
		dayIndex := day + 1
		fmt.Print(dayIndex, ",")

		for i, numStreaks := range dayStreaks {
			fmt.Print(numStreaks)

			if i < len(dayStreaks)-1 {
				fmt.Print(",")
			} else {
				fmt.Println()
			}
		}
	}
}
