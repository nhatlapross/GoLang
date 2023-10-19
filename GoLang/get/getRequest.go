package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcom to web")
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
