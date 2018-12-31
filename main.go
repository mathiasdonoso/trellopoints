package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adlio/trello"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	boardID := "rQVhgPGJ"

	client := trello.NewClient(os.Getenv("TRELLO_KEY"), os.Getenv("TRELLO_TOKEN"))
	board, err := client.GetBoard(boardID, trello.Defaults())
	handleError(err)

	lists, err := board.GetLists(trello.Defaults())
	handleError(err)

	for _, list := range lists {
		points := storyPoints(list)
		args := make(map[string]string)
		args["value"] = fmt.Sprintf("%s (%.1f)", NameWithoutPoints(list.Name), points)

		client.Put(fmt.Sprintf("/list/%s/name", list.ID), args, list)
	}

	w.Header().Set("Content-Type", "application/json ")
	w.Write([]byte("200 - OK"))
}

func storyPoints(list *trello.List) float64 {
	var listStoryPoints float64
	cards, err := list.GetCards(trello.Defaults())
	handleError(err)

	for _, card := range cards {
		if HasStoryPoint(card.Name) {
			points, err := GetStoryPoint(card.Name)
			handleError(err)
			listStoryPoints += points
		}
	}

	return listStoryPoints
}
