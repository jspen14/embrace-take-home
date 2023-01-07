package csv_retriever

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"embrace.io/take_home/internal/types"
)

var exists void

func (r csvRetriever) GetDaysActiveByUserID() (types.DaysActiveByUserID, error) {

	// the value is a set
	s := make(map[types.UserID]types.DaysActiveSet)

	f, err := os.Open(r.dataFile)
	if err != nil {
		return s, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return s, err
		}

		timestamp, err := strconv.ParseInt(rec[0], 10, 32)
		if err != nil {
			return s, err
		}

		dt := time.Unix(timestamp, 0).UTC()
		day := dt.Day()

		userIDStr, err := strconv.ParseInt(rec[1], 10, 64)
		if err != nil {
			return s, err
		}

		userID := types.UserID(userIDStr)

		//  this uses compression by only storing one element per day.
		// 	from my tests, this was able to considerably reduce the amount of data being processed.
		if _, ok := s[userID]; ok {
			s[userID][day] = exists
		} else {
			s[userID] = make(map[int]struct{})
			s[userID][day] = exists
		}
	}

	return s, nil
}
