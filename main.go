package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zuckeyM-17/vocabulary-notion/deepl"
	"github.com/zuckeyM-17/vocabulary-notion/notion"

	"github.com/joho/godotenv"
)

var (
	errLog = log.New(os.Stderr, "[Error] ", 0)
)

func main() {
	flag.Parse()
	var (
		args   = flag.Args()
		wordEn = args[0]
	)

	err := godotenv.Load()
	if err != nil {
		errLog.Println("Error loading .env file")
	}

	notionToken := os.Getenv("NOTION_TOKEN")
	deeplApiKey := os.Getenv("DEEPL_API_KEY")
	databaseId := os.Getenv("DATABASE_ID")

	wordEntry, err := notion.SearchWord(wordEn, notionToken, databaseId)
	if err != nil {
		errLog.Println(err)
	}

	wordJa := ""
	if wordEntry.Id == "" {
		w, err := deepl.Translate(wordEn, deeplApiKey)
		wordJa = w
		if err != nil {
			errLog.Println(err)
		}

		notion.InsertWord(wordEn, wordJa, notionToken, databaseId)
	} else {
		wordJa = wordEntry.WordJa
		notion.IncrementCount(wordEntry.Id, wordEntry.Count, notionToken)
	}

	fmt.Println(wordEn, wordJa)
}
