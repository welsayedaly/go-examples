package main

import (
	"fmt"
	"log"
	"time"
)

// Format Date
func formatDate() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
}

//Parse Date
func parseDate() {
	const YYYYMMDD = "2006-01-01"

	s := "2022-02-02"
	x, err := time.Parse(YYYYMMDD, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)

}

// Parse Date with Time Zone
func parseDateWithTimeZone() {
	s := "2022-03-23T07:00:00+01:00"
	loc, _ := time.LoadLocation("Europe/Berlin")
	t, err := time.ParseInLocation(time.RFC3339, s, loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
}
func mainD() {
	gitBranch()
	formatDate()
	parseDate()
	parseDateWithTimeZone()

}

/**

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)


*/
