package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// Get Epoch timestamp.
	ts := now.Unix()
	fmt.Println(ts)

	// Get time zone and offset.
	zone, offset := now.Zone()
	fmt.Println(zone, offset)

	// Get formatted date and time.
	fmt.Println(now.Format(time.RFC3339))

	// Set time zone
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	fmt.Println(now.In(loc).Format(time.RFC3339))

	// Datetime diff
	time.Sleep(time.Second)
	passOneSecond := time.Now()
	fmt.Println("Time diffrence:", passOneSecond.Sub(now))

	// Timsstamp to datetime
	dt := time.Unix(ts, 0)
	fmt.Println(dt)
}
