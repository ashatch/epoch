package main

import "fmt"

func main() {
	flags := GetEpochFlags()
	currentTime := GetTime(flags.utc)

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

func iso(time TimeData) {
	fmt.Printf("%s", time.iso)
}

func seconds(time TimeData) {
	fmt.Printf("%d", time.seconds)
}

func millis(time TimeData) {
	fmt.Printf("%d", time.millis)
}

func nanos(time TimeData) {
	fmt.Printf("%d", time.nanos)
}

func fullOutput(time TimeData) {
	fmt.Printf("%8s: %s\n", "Zone", time.zone)
	fmt.Printf("%8s: %s\n", "Now", time.iso)
	fmt.Printf("%8s: %d\n", "Seconds", time.seconds)
	fmt.Printf("%8s: %d\n", "Millis", time.millis)
	fmt.Printf("%8s: %d\n", "Nanos", time.nanos)
}
