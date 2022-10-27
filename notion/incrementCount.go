package notion

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func IncrementCount(pageId string, count int64, notionToken string) {
	count = count + 1

	type Count struct {
		Number int64 `json:"number"`
	}

	type Properties struct {
		Count Count `json:"Count"`
	}

	type Payload struct {
		Properties Properties `json:"properties"`
	}

	data := Payload{
		Properties: Properties{
			Count: Count{
				Number: count,
			},
		},
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	req, err := http.NewRequest("PATCH", "https://api.notion.com/v1/pages/"+pageId, bytes.NewReader(payloadBytes))

	if err != nil {
		// handle err
	}

	req.Header.Set("Authorization", "Bearer "+notionToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}

	defer resp.Body.Close()
}
