package main

import "os"

var (
	firesword_log chan string
)

func init() {
	firesword_log = make(chan string)
}

func StartFSLog() {
	if f_firesword_log == "" {
		return
	}

	f, err := os.OpenFile(f_firesword_log, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for {
		fsl := <-firesword_log

		if _, err = f.WriteString(fsl + "\n"); err != nil {
			panic(err)
		}
	}

}
