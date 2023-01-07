package aggregator

import "embrace.io/take_home/internal/types"

type Aggregator interface {
	ProcessData(data types.DaysActiveByUserID) error
	GetResults() types.StreakResults
	PrintResults()
}
