package v1_aggregator

import (
	"embrace.io/take_home/internal/aggregator"
	"embrace.io/take_home/internal/types"
)

const (
	startDay = 1
	endDay   = 14
)

type v1Aggregator struct {
	aggregates types.StreakResults
}

func NewV1Aggregator() (aggregator.Aggregator, error) {
	return &v1Aggregator{}, nil
}
