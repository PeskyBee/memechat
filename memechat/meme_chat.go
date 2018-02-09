package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const memeSearchEndpoint = "https://imgflip.com/memesearch?q="
const memePageEndpoint = "https://imgflip.com"

func getAttr(query string, selector string, attr string) (string, error) {
	doc, err := goquery.NewDocument(query)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Find the review items
	imageRes := doc.Find(selector).First()
	memeURL, ok := imageRes.Attr(attr)
	if ok != true {
		err := fmt.Errorf("Could not find attr %v for query: %v", attr, query)
		log.Fatal(err)
		return "", err
	}
	return memeURL, nil
}

func getMemePageURL(memeName string) (string, error) {
	queryParam := url.QueryEscape(memeName)
	query := memeSearchEndpoint + queryParam
	return getAttr(query, "#memeTemplates .mt-box .mt-caption", "href")
}

func getMemeImageURL(memeURL string) string {
	// memeUrl normally contains "/memegenerator/<meme_name>", extract meme_name
	arrURL := strings.Split(memeURL, "/")
	return arrURL[len(arrURL)-1]
}

// func GetImageUrlForMeme(memeName string) (string, error) {
// 	memeUrl, err := getMemePageUrl(memeName)
// 	if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	}
// 	memeImageUrl, err := getMemeImageUrl(memeUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 		return "", err
// 	}
// 	return memeImageUrl, nil
// }

// func main() {
// 	memeName := os.Args[1]
// 	log.Printf("search for %v", memeName)
// 	// memeUrl, err := GetImageUrlForMeme(memeName)
// 	// log.Println(memeUrl)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return
// 	// }
// }
