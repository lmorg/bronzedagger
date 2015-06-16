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
	if f_debug {
		log.Println(v...)
	}
}
