package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup()
	logfile, logerr := os.OpenFile("filezilla.log", os.O_RDONLY, 0666)
	if logerr != nil {
		fmt.Println(logerr)
	}
	buff := make([]byte, 1024)
	for n, logerr := logfile.Read(buff); logerr == nil; n, logerr = logfile.Read(buff) {
		// go fmt.Print(string(buff[:n]))

	}
	if logerr != nil {
		panic(fmt.Sprintf("Read occurs error: %s", logerr))
	}
}

func ShowMessage(wg *sync.WaitGroup, buff string) {
	fmt.Println(buff)
}
