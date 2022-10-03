package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RSS struct
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

//Channel struct
type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

//Item struct
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

//News struct
type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	
	//data := readGoogleTrends(getGoogleTrends())
	data := readGoogleTrends()
	//err := xml.Unmarshal([]byte(data), &r)
	err := xml.Unmarshal(data, &r)

	// handle a parsing error
	if err != nil {
		fmt.Println("error:", err)
	}


	fmt.Println("\nHere are the Google search trends for today")
	fmt.Println("-------------------------------------------")

	// print out the results
	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link to trend:", r.Channel.ItemList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("Link to article:", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("--------------------------")
	}

}

// go and pull the xml from the google trends website
// and return the response
func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

// read the http.response and convert it into a []byte
func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
