package v1_aggregator

import (
	"embrace.io/take_home/internal/types"
)

func (a *v1Aggregator) ProcessData(data types.DaysActiveByUserID) error {

	for currDay := startDay; currDay <= endDay; currDay++ {
		prevDay := currDay - 1

		dayStreaks := make([]int, 14)

		for _, daysActive := range data {

			_, prevDayExists := daysActive[prevDay]
			_, todayExists := daysActive[currDay]

			wasStreakStartedToday := !prevDayExists && todayExists

			if wasStreakStartedToday {
				//  Figure out how long the streak was
				streakLength := 1
				nextDay := currDay + 1

				for {
					_, nextDayExists := daysActive[nextDay]

					if nextDayExists {
						streakLength += 1
						nextDay += 1
					} else {
						break
					}
				}

				dayStreaks[streakLength-1] += 1
			}
		}

		a.aggregates = append(a.aggregates, dayStreaks)
	}

	return nil
}
