package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`

type Message struct {
	Name string
	Text string
}

func main() {
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonStream))
	var m Message
	for {
		if err := jsonDecoder.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", m)
	}
}
