package notion

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

var (
	errLog = log.New(os.Stderr, "[Error] ", 0)
)

type WordEntry struct {
	Id     string
	WordEn string
	WordJa string
	Count  int64
}

func SearchWord(wordEn string, notionToken string, databaseId string) (WordEntry, error) {
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

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(d))
	if err != nil {
		errLog.Println(err)
		return WordEntry{}, err
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
		return WordEntry{}, err
	}
	defer res.Body.Close()

	r, _ := io.ReadAll(res.Body)
	json := string(r)

	var (
		id     = gjson.Get(json, "results.0.id").String()
		wordJa = gjson.Get(json, "results.0.properties.Japanese.rich_text.0.text.content").String()
		count  = gjson.Get(json, "results.0.properties.Count.number").Int()
	)
	return WordEntry{
		Id:     id,
		WordJa: wordJa,
		WordEn: wordEn,
		Count:  count,
	}, nil
}
