package main

import (
	"fmt"

	"github.com/nlopes/slack"
	"os"
)

func main() {
	token := os.Getenv("GREMLIN_SLACK_API_TOKEN")
	api := slack.New(token)
	params := slack.NewSearchParameters()

	msgs, err := api.SearchMessages("from:@Daniel on:Today", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message Count: %d\n", msgs.Total)
}
