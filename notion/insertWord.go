package notion

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/zuckeyM-17/vocabulary-notion/util"
)

func InsertWord(wordEn string, wordJa string, notionToken string, databaseId string) {

	type Text struct {
		Content string `json:"content"`
	}
	type Title struct {
		Text Text `json:"text"`
	}

	type English struct {
		Title []Title `json:"title"`
	}
	type RichText struct {
		Text Text `json:"text"`
	}
	type Japanese struct {
		RichText []RichText `json:"rich_text"`
	}
	type Count struct {
		Number int `json:"number"`
	}

	type Parent struct {
		DatabaseID string `json:"database_id"`
	}
	type Properties struct {
		English  English  `json:"English"`
		Japanese Japanese `json:"Japanese"`
		Count    Count    `json:"Count"`
	}

	type Payload struct {
		Parent     Parent     `json:"parent"`
		Properties Properties `json:"properties"`
	}

	data := Payload{
		Parent: Parent{DatabaseID: databaseId},
		Properties: Properties{
			English:  English{Title: []Title{{Text{Content: wordEn}}}},
			Japanese: Japanese{RichText: []RichText{{Text{Content: wordJa}}}},
			Count:    Count{Number: 0},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		util.ErrLog(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.notion.com/v1/pages", body)
	if err != nil {
		util.ErrLog(err)
	}
	req.Header.Set("Authorization", "Bearer "+notionToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2021-08-16")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.ErrLog(err)
	}
	defer resp.Body.Close()
}
