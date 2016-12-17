package main

import (
	"log"
	"os"
)

func isErr(err error) bool {
	if err != nil {
		log.Println("ERROR!!", err)
		os.Exit(254)
	}
	return false
}

func debugLog(v ...interface{}) {
	if fDebug {
		log.Println(v...)
	}
}
