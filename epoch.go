package main

import "fmt"
import flags "github.com/ashatch/epoch/flags"
import timeinfo "github.com/ashatch/epoch/timeinfo"

func main() {
	flags := flags.GetEpochFlags()
	currentTime := timeinfo.GetTime(flags.Utc)

	if flags.Iso {
		iso(currentTime)
	} else if flags.Seconds {
		seconds(currentTime)
	} else if flags.Millis {
		millis(currentTime)
	} else if flags.Nanos {
		nanos(currentTime)
	} else {
		fullOutput(currentTime)
	}
}

func iso(time timeinfo.TimeData) {
	fmt.Printf("%s", time.Iso)
}

func seconds(time timeinfo.TimeData) {
	fmt.Printf("%d", time.Seconds)
}

func millis(time timeinfo.TimeData) {
	fmt.Printf("%d", time.Millis)
}

func nanos(time timeinfo.TimeData) {
	fmt.Printf("%d", time.Nanos)
}

func fullOutput(time timeinfo.TimeData) {
	fmt.Printf("%8s: %s\n", "Zone", time.Zone)
	fmt.Printf("%8s: %s\n", "Now", time.Iso)
	fmt.Printf("%8s: %d\n", "Seconds", time.Seconds)
	fmt.Printf("%8s: %d\n", "Millis", time.Millis)
	fmt.Printf("%8s: %d\n", "Nanos", time.Nanos)
}
