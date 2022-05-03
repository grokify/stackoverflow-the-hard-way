package main

import (
	"fmt"
	"log"

	stackoverflow "github.com/grokify/go-stackoverflow/client"
	"github.com/grokify/go-stackoverflow/util"
)

func main() {
	apiClient := stackoverflow.NewAPIClient(stackoverflow.NewConfiguration())
	history, err := util.GetReputationHistoryAll(
		apiClient, "stackoverflow", "1908967")
	if err != nil {
		log.Fatal(err)
	}

	// history.DateForReputation calculate the dates when a certain reputation
	// score was achieved.
	dayRep, err := history.DateForReputation(int32(2000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DATE [%v]\n", dayRep.Day)
}
