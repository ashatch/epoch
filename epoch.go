package main

import "fmt"
import "time"
import "flag"

type epochFlags struct {
	utc     bool
	iso     bool
	seconds bool
	millis  bool
	nanos   bool
}

type timeData struct {
	iso     string
	zone    string
	seconds int64
	millis  int64
	nanos   int64
}

func main() {
	flags := epochFlags{}

	flag.BoolVar(&flags.utc, "utc", false, "Use time in UTC")
	flag.BoolVar(&flags.iso, "iso", false, "Show ISO8601/RFC3339 format only")
	flag.BoolVar(&flags.seconds, "seconds", false, "Show seconds only")
	flag.BoolVar(&flags.millis, "millis", false, "Show milliseconds only")
	flag.BoolVar(&flags.nanos, "nanos", false, "Show nanoseconds only")
	flag.Parse()

	var now = timeForNow(flags.utc)
	zoneName, _ := now.Zone()

	currentTime := timeData{
		zone:    zoneName,
		iso:     now.Format(time.RFC3339),
		seconds: now.Unix(),
		millis:  now.UnixNano() / 1000000,
		nanos:   now.UnixNano()}

	if flags.iso {
		iso(currentTime)
	} else if flags.seconds {
		seconds(currentTime)
	} else if flags.millis {
		millis(currentTime)
	} else if flags.nanos {
		nanos(currentTime)
	} else {
		fullOutput(currentTime)
	}
}

func timeForNow(useUtc bool) time.Time {
	var now = time.Now()
	if useUtc {
		utcLocation, _ := time.LoadLocation("UTC")
		now = time.Now().In(utcLocation)
	}
	return now
}

func iso(time timeData) {
	fmt.Printf("%s", time.iso)
}

func seconds(time timeData) {
	fmt.Printf("%d", time.seconds)
}

func millis(time timeData) {
	fmt.Printf("%d", time.millis)
}

func nanos(time timeData) {
	fmt.Printf("%d", time.nanos)
}

func fullOutput(time timeData) {
	fmt.Printf("%8s: %s\n", "Zone", time.zone)
	fmt.Printf("%8s: %s\n", "Now", time.iso)
	fmt.Printf("%8s: %d\n", "Seconds", time.seconds)
	fmt.Printf("%8s: %d\n", "Millis", time.millis)
	fmt.Printf("%8s: %d\n", "Nanos", time.nanos)
}
