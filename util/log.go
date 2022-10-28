package util

import (
	"log"
	"os"
)

var (
	errLog = log.New(os.Stderr, "Error ", 1)
)

func ErrLog(e error) {
	errLog.Println(e)
}
