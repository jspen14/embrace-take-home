package csv_retriever_test

import (
	"testing"

	csvRetriever "embrace.io/take_home/internal/retriever/csv"
	"github.com/stretchr/testify/assert"
)

func TestCSVReceiver(t *testing.T) {
	t.Run("Happy path - receives all data from file", func(t *testing.T) {
		retriever, err := csvRetriever.NewCSVRetriever("./test_data/test_data.csv")
		assert.Nil(t, err)

		data, err := retriever.GetDaysActiveByUserID()
		assert.Nil(t, err)

		assert.Equal(t, 5, len(data))

		totalCount := 0

		for _, sessionDays := range data {
			totalCount += len(sessionDays)
		}

		assert.Equal(t, 13, totalCount)

		user1SessionDays := data[1]

		assert.Equal(t, 5, len(user1SessionDays))
	})
}
