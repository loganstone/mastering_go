package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	name := filepath.Base(os.Args[0])

	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, name)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go.")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some Application")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Logging in Go.")
	fmt.Println("Will you see this?")
}
