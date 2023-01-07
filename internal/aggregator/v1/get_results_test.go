package v1_aggregator_test

import (
	"testing"

	v1Aggregator "embrace.io/take_home/internal/aggregator/v1"
	"github.com/stretchr/testify/assert"
)

func TestV1Aggregator_GetResults(t *testing.T) {
	t.Run("Happy path - gets data", func(t *testing.T) {
		aggregator, err := v1Aggregator.NewV1Aggregator()
		assert.Nil(t, err)

		results := aggregator.GetResults()

		emptyRow := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 1; i <= len(results); i++ {

			assert.Equal(t, emptyRow, results[i])
		}
	})
}
