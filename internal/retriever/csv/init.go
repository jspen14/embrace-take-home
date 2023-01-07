package csv_retriever

import (
	"embrace.io/take_home/internal/retriever"
)

type void struct{}

type csvRetriever struct {
	dataFile string
}

func NewCSVRetriever(dataFile string) (retriever.Retriever, error) {

	return csvRetriever{
		dataFile: dataFile,
	}, nil
}
