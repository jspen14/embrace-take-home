package retriever

import "embrace.io/take_home/internal/types"

type Retriever interface {
	GetDaysActiveByUserID() (types.DaysActiveByUserID, error)
}
