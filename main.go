package main

import (
	"os"

	"github.com/adlio/trello"
)

func main() {
	client := trello.NewClient(os.Getenv("TRELLO_KEY"), os.Getenv("TRELLO_TOKEN"))
	client.GetBoard("7FAXeNuf", trello.Defaults())
}
