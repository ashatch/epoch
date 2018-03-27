package main

import "time"

// TimeData contains epoch information
type TimeData struct {
	iso     string
	zone    string
	seconds int64
	millis  int64
	nanos   int64
}

// GetTime retrieves a TimeData
func GetTime(useUTC bool) TimeData {

	var now = timeForNow(useUTC)
	zoneName, _ := now.Zone()

	currentTime := TimeData{
		zone:    zoneName,
		iso:     now.Format(time.RFC3339),
		seconds: now.Unix(),
		millis:  now.UnixNano() / 1000000,
		nanos:   now.UnixNano()}

	return currentTime
}

func timeForNow(useUtc bool) time.Time {
	var now = time.Now()
	if useUtc {
		utcLocation, _ := time.LoadLocation("UTC")
		now = time.Now().In(utcLocation)
	}
	return now
}
