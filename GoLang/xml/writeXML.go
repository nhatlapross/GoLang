package main

import (
	"encoding/xml"
	"io/ioutil"
)

type notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    body   `xml:"body"`
}

type body struct {
	Text1 string `xml:"text1"`
	Text2 string `xml:"text2"`
}

func main() {
	body := body{
		Text1: "Hello",
		Text2: "World",
	}

	note := notes{
		To:      "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    body,
	}

	file, _ := xml.MarshalIndent(note, "", "    ")

	ioutil.WriteFile("notes1.xml", file, 0644)

}
