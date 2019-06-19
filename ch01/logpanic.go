package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_ALERT|syslog.LOG_MAIL, "logpanic.go")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}
	// output with error information, exit status 2
	log.Panic(sysLog)
	fmt.Println("Will you see this?")
}
