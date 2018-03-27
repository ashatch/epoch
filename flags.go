package main

import "flag"

// EpochFlags represents flags given to epoch program
type EpochFlags struct {
	utc     bool
	iso     bool
	seconds bool
	millis  bool
	nanos   bool
}

// GetEpochFlags gives an EpochFlags based on command line args
func GetEpochFlags() EpochFlags {
	flags := EpochFlags{}

	flag.BoolVar(&flags.utc, "utc", false, "Use time in UTC")
	flag.BoolVar(&flags.iso, "iso", false, "Show ISO8601/RFC3339 format only")
	flag.BoolVar(&flags.seconds, "seconds", false, "Show seconds only")
	flag.BoolVar(&flags.millis, "millis", false, "Show milliseconds only")
	flag.BoolVar(&flags.nanos, "nanos", false, "Show nanoseconds only")
	flag.Parse()

	return flags
}
