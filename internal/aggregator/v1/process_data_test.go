package v1_aggregator_test

import (
	"testing"

	v1Aggregator "embrace.io/take_home/internal/aggregator/v1"
	"embrace.io/take_home/internal/types"
	"github.com/stretchr/testify/assert"
)

type void struct{}

var exists void

func TestV1Aggregator_ProcessData(t *testing.T) {
	t.Run("Happy path - computes data", func(t *testing.T) {
		data := make(types.DaysActiveByUserID)
		data[1] = make(map[int]struct{})
		data[1][1] = exists
		data[1][2] = exists
		data[1][3] = exists
		data[1][4] = exists
		data[1][5] = exists

		data[2] = make(map[int]struct{})
		data[2][1] = exists
		data[2][3] = exists
		data[2][4] = exists

		data[3] = make(map[int]struct{})
		data[3][1] = exists
		data[3][2] = exists
		data[3][3] = exists

		data[4] = make(map[int]struct{})
		data[4][1] = exists

		data[5] = make(map[int]struct{})
		data[5][5] = exists

		aggregator, err := v1Aggregator.NewV1Aggregator()
		assert.Nil(t, err)

		aggregator.ProcessData(data)

		results := aggregator.GetResults()

		assert.Equal(t, 14, len(results))
		assert.Equal(t, 14, len(results[0]))

		row1 := []int{2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row3 := []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row4 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row5 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row6 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row7 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row8 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row9 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row10 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row11 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row12 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row13 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		row14 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		assert.Equal(t, row1, results[0])
		assert.Equal(t, row2, results[1])
		assert.Equal(t, row3, results[2])
		assert.Equal(t, row4, results[3])
		assert.Equal(t, row5, results[4])
		assert.Equal(t, row6, results[5])
		assert.Equal(t, row7, results[6])
		assert.Equal(t, row8, results[7])
		assert.Equal(t, row9, results[8])
		assert.Equal(t, row10, results[9])
		assert.Equal(t, row11, results[10])
		assert.Equal(t, row12, results[11])
		assert.Equal(t, row13, results[12])
		assert.Equal(t, row14, results[13])
	})
}
