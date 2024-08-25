package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"os"
	"github.com/go-git/go-git/v5"
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

func errCheck(err error) {
	if err != nil {
		msg := "Error! " + err.Error()
		panic(msg)
	}
}

func fetchAndParseRSS(url string) (*RSS, error){

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


func main() {

	url := "http://aur.archlinux.org/rss"
	rss,err := fetchAndParseRSS(url)
	errCheck(err)

	fmt.Printf("Channel: %s\nDescription: %s\n\n", rss.Channel.Title, rss.Channel.Description)
	for _, item := range rss.Channel.Items {
		fmt.Printf("Title: %s\nLink: %s\nPublished Date: %s\nDescription: %s\n\n", item.Title, item.Link, item.PubData, item.Description)
		//repo := strings.Replace(s string, old string, new string, n int)
		repo := strings.Replace(item.Link, "packages/", "", 1)

		// TODO: check the git clone, and send a message to a repo or something.
		_, _:= git.PlainClone(item.Title + ".git", false, &git.CloneOptions{
			URL: repo,
			Progress: os.Stdout,
		})
	}
}
