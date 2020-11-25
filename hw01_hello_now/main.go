package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const Layout = "2006-01-02 15:04:05 -0700 MST"

func Today() string {
	return time.Now().Round(0).Format(Layout)
}
func TodayNtp() string {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Fatal Error: %v\n", err)
	}
	return ntpTime.Round(0).Format(Layout)
}
func main() {
	fmt.Println("current time:", Today())
	fmt.Println("exact time:", TodayNtp())
}
