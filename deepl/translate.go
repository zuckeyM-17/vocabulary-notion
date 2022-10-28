package deepl

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/zuckeyM-17/vocabulary-notion/util"
)

func Translate(wordEn string, deeplApiKey string) (string, error) {
	var (
		authorization = "DeepL-Auth-Key " + deeplApiKey
		uri           = "https://api-free.deepl.com/v2/translate"
	)

	params := url.Values{}
	params.Add("text", wordEn)
	params.Add("target_lang", `JA`)

	req, err := http.NewRequest("POST", uri, strings.NewReader(params.Encode()))
	if err != nil {
		util.ErrLog(err)
		return "", err
	}

	req.Header.Add("Authorization", authorization)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		util.ErrLog(err)
		return "", err
	}

	defer res.Body.Close()

	r, _ := io.ReadAll(res.Body)
	json := string(r)

	wordJa := gjson.Get(json, "translations.0.text").String()

	return wordJa, nil
}
