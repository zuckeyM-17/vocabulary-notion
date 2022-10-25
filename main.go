package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

var (
	errLog = log.New(os.Stderr, "[Error] ", 0)
)

func searchWord(wordEn string, notionToken string, databaseId string) (string, string, int64) {
	var (
		uri           = "https://api.notion.com/v1/databases/" + databaseId + "/query"
		authorization = "Bearer " + notionToken
		contentType   = "application/json"
		notionVersion = "2022-06-28"
	)

	type TitleFilterData struct {
		Equals string `json:"equals"`
	}

	type FilterData struct {
		Property string          `json:"property"`
		Title    TitleFilterData `json:"title"`
	}

	type SearchData struct {
		Filter FilterData `json:"filter"`
	}

	searchData := SearchData{
		Filter: FilterData{
			Property: "English",
			Title: TitleFilterData{
				Equals: wordEn,
			},
		},
	}

	d, _ := json.Marshal(searchData)

	fmt.Println(string(d))

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(d))
	if err != nil {
		errLog.Println(err)
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Notion-Version", notionVersion)

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		errLog.Println(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)

	r, _ := io.ReadAll(res.Body)
	json := string(r)

	var (
		id     = gjson.Get(json, "results.0.id").String()
		wordJa = gjson.Get(json, "results.0.properties.Japanese.rich_text.0.text.content").String()
		count  = gjson.Get(json, "results.0.properties.Count.number").Int()
	)
	return id, wordJa, count
}

func main() {
	flag.Parse()
	var (
		args   = flag.Args()
		wordEn = args[0]
	)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	notionToken := os.Getenv("NOTION_TOKEN")
	// deeplApiKey := os.Getenv("DEEPL_API_KEY")
	databaseId := os.Getenv("DATABASE_ID")

	id, wordJa, count := searchWord(wordEn, notionToken, databaseId)
	fmt.Println(id, wordJa, count)
}
