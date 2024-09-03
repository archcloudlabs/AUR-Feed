package main


import (
	"encoding/xml"
	"net/http"
	"io"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}


type Channel struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Link string `xml:"link"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubData string `xml:"PubData"`
}

func FetchAndParseRSS(url string) (*RSS, error){

	resp, err := http.Get(url)
	errCheck(err)

	defer resp.Body.Close() // close the handle later.

	data, err := io.ReadAll(resp.Body)
	errCheck(err)

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	errCheck(err)

	return &rss, nil
}
