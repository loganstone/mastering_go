package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	// backticks is raw string literal, do not need escaped
	// https://golang.org/ref/spec#String_literals
	regexStrToDt1 = `.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`
	regexStrToDt2 = `.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`
	timeParseStr1 = "02/Jan/2006:15:04:05 -0700"
	timeParseStr2 = "Jan-02-06:15:04:05 -0700"
)

var (
	r1 = regexp.MustCompile(regexStrToDt1)
	r2 = regexp.MustCompile(regexStrToDt2)
)

func findAndReplace(r *regexp.Regexp, p string, l string) (string, error) {
	m := r.FindStringSubmatch(l)
	d, err := time.Parse(p, m[1])
	if err != nil {
		return "", err
	}
	newFmt := d.Format(time.Stamp)
	// func Replace(s, old, new string, n int) string
	// if n < 0, there is no limit on the number of replacements
	// that mean is "replace all"
	return strings.Replace(l, m[1], newFmt, 1), nil
}

func main() {

	if len(os.Args) == 1 {
		log.Fatalln("Please provide one text file to process!")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("error opening file %s", err)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}

		if r1.MatchString(line) {
			line, err := findAndReplace(r1, timeParseStr1, line)
			if err == nil {
				fmt.Print(line)
			} else {
				notAMatch++
			}
			continue
		}

		if r2.MatchString(line) {
			line, err := findAndReplace(r2, timeParseStr2, line)
			if err == nil {
				fmt.Print(line)
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match!")
}
