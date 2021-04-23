// https://tutorialedge.net/golang/consuming-restful-api-with-go/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Results struct {
	News []News `json:"results"`
}

type News struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func main() {
	resp, err := http.Get("https://actu.epfl.ch/api/v1/news/")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	//anyBody := string(anyBody.results)
	//var myString strings.Builder

	var resultObject Results
	json.Unmarshal(body, &resultObject)

	// https://www.geeksforgeeks.org/different-ways-to-concatenate-two-strings-in-golang/
	for i := 0; i < len(resultObject.News); i++ {
		title := string(resultObject.News[i].Title)
		subtitle := string(resultObject.News[i].Subtitle)
		log.Println(title + "\n" + subtitle  +  "\n\n")
	}

}
