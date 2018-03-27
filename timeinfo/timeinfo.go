package timeinfo

import "time"

// TimeData contains epoch information
type TimeData struct {
	Iso     string
	Zone    string
	Seconds int64
	Millis  int64
	Nanos   int64
}

// GetTime retrieves a TimeData based on current system time or with useUTC true for UTC
func GetTime(useUTC bool) TimeData {

	var now = timeForNow(useUTC)
	zoneName, _ := now.Zone()

	currentTime := TimeData{
		Zone:    zoneName,
		Iso:     now.Format(time.RFC3339),
		Seconds: now.Unix(),
		Millis:  now.UnixNano() / 1000000,
		Nanos:   now.UnixNano()}

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
