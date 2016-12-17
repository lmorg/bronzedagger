package main

import "os"

var (
	fsLog chan string
)

func init() {
	fsLog = make(chan string)
}

func StartFSLog() {
	if fFsLog == "" {
		return
	}

	f, err := os.OpenFile(fFsLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for {
		fsl := <-fsLog

		if _, err = f.WriteString(fsl + "\n"); err != nil {
			panic(err)
		}
	}

}
