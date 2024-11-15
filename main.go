package main

import (
	"fmt"
	"os"
	"strings"

	aur "github.com/AUR-Feed/pkg"
	"github.com/go-git/go-git/v5"
)


func main() {

	url := "http://aur.archlinux.org/rss"
	rss, err := aur.FetchAndParseRSS(url)
	aur.ErrCheck(err)

	fmt.Printf("Channel: %s\nDescription: %s\n\n", rss.Channel.Title, rss.Channel.Description)
	for _, item := range rss.Channel.Items {
		fmt.Printf("Title: %s\nLink: %s\nPublished Date: %s\nDescription: %s\n\n", item.Title, item.Link, item.PubData, item.Description)
		//repo := strings.Replace(s string, old string, new string, n int)
		repo := strings.Replace(item.Link, "packages/", "", 1)

		// TODO: check the git clone, and send a message to a repo or something.
		_, _ = git.PlainClone(item.Title+".git", false, &git.CloneOptions{
			URL:      repo,
			Progress: os.Stdout,
		})
	}
}
