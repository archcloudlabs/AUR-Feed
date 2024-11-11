package aur


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

func ErrCheck(err error) {
	if err != nil {
		msg := "Error! " + err.Error()
		panic(msg)
	}
}

func FetchAndParseRSS(url string) (*RSS, error){

	resp, err := http.Get(url)
	ErrCheck(err)

	defer resp.Body.Close() // close the handle later.

	data, err := io.ReadAll(resp.Body)
	ErrCheck(err)

	var rss RSS
	err = xml.Unmarshal(data, &rss)
	ErrCheck(err)

	return &rss, nil
}
