package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT, "logfatal.go")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}
	// print log and exit program. exit status 1
	log.Fatal(sysLog)
	fmt.Println("Will you see this?")
}
