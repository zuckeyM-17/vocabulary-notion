package deepl

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

var (
	errLog = log.New(os.Stderr, "[Error] ", 0)
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
		errLog.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", authorization)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		errLog.Println(err)
		return "", err
	}

	defer res.Body.Close()

	r, _ := io.ReadAll(res.Body)
	json := string(r)

	wordJa := gjson.Get(json, "translations.0.text").String()

	return wordJa, nil
}
