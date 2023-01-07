package v1_aggregator

import "embrace.io/take_home/internal/types"

func (a *v1Aggregator) GetResults() types.StreakResults {
	return a.aggregates
}
