package main

import (
	"log"
	"os"

	v1Aggregator "embrace.io/take_home/internal/aggregator/v1"
	csvRetriever "embrace.io/take_home/internal/retriever/csv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("valid usage: go run main.go pathToFile")
	}

	dataFile := os.Args[1]

	retriever, err := csvRetriever.NewCSVRetriever(dataFile)
	if err != nil {
		log.Fatal("unable to create data retriever")
	}

	data, err := retriever.GetDaysActiveByUserID()
	if err != nil {
		log.Fatal("unable to retrieve data")
	}

	aggregator, err := v1Aggregator.NewV1Aggregator()
	if err != nil {
		log.Fatal("unable to create aggregator")
	}

	err = aggregator.ProcessData(data)
	if err != nil {
		log.Fatal("unable to process data")
	}

	aggregator.PrintResults()
}
