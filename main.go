package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	waitUntil := os.Args[1]
	relativeTo := now()
	waitUntilTime, duration := calc(waitUntil, relativeTo)

	fmt.Printf("Now:        %v\n", relativeTo)
	fmt.Printf("Wait until: %v\n", waitUntilTime)
	fmt.Printf("Duration:   %v\n", duration)
	fmt.Printf("\n")
	time.Sleep(duration)
}

func now() time.Time {
	var now = time.Now().Local()
	now = now.Add(-time.Duration(now.Nanosecond()) * time.Nanosecond)
	return now
}

func calc(waitUntil string, relativeTo time.Time) (time.Time, time.Duration) {
	waitUntilParts := strings.Split(waitUntil, ":")
	waitUntilTime := relativeTo
	switch len(waitUntilParts) {
	case 1:
		h := atoi(waitUntilParts[0])
		if h < waitUntilTime.Hour() || h == waitUntilTime.Hour() && waitUntilTime.Minute() > 0 {
			// h=15, now=16:23 => wait until tomorrow
			waitUntilTime = waitUntilTime.Add(time.Duration(h-waitUntilTime.Hour()+24) * time.Hour)
		} else {
			// h=17, now=16:23 => wait until full hour
			waitUntilTime = waitUntilTime.Add(time.Duration(h-waitUntilTime.Hour()) * time.Hour)
		}
		// clear minutes and seconds
		waitUntilTime = waitUntilTime.Add(-time.Duration(waitUntilTime.Minute()) * time.Minute)
		waitUntilTime = waitUntilTime.Add(-time.Duration(waitUntilTime.Second()) * time.Second)
	case 2:
		h := atoi(waitUntilParts[0])
		m := atoi(waitUntilParts[1])
		if h < waitUntilTime.Hour() || h == waitUntilTime.Hour() && waitUntilTime.Minute() > m {
			// h=15, m=30, now=16:23
			waitUntilTime = waitUntilTime.Add(time.Duration(h-waitUntilTime.Hour()+24) * time.Hour)
			waitUntilTime = waitUntilTime.Add(time.Duration(m-waitUntilTime.Minute()) * time.Minute)
		} else {
			// h=17, m=30, now=16:23
			waitUntilTime = waitUntilTime.Add(time.Duration(h-waitUntilTime.Hour()) * time.Hour)
			waitUntilTime = waitUntilTime.Add(time.Duration(m-waitUntilTime.Minute()) * time.Minute)
		}
		// clear seconds
		waitUntilTime = waitUntilTime.Add(-time.Duration(waitUntilTime.Second()) * time.Second)
	default:
		panic(fmt.Errorf("unsupported time format: %v", waitUntil))
	}

	duration := waitUntilTime.Sub(relativeTo)
	duration -= duration % time.Second

	return waitUntilTime, duration
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
