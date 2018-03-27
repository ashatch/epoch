package flags

import "flag"

// EpochFlags represents flags given to epoch program
type EpochFlags struct {
	Utc     bool
	Iso     bool
	Seconds bool
	Millis  bool
	Nanos   bool
}

// GetEpochFlags gives an EpochFlags based on command line args
func GetEpochFlags() EpochFlags {
	flags := EpochFlags{}

	flag.BoolVar(&flags.Utc, "utc", false, "Use time in UTC")
	flag.BoolVar(&flags.Iso, "iso", false, "Show ISO8601/RFC3339 format only")
	flag.BoolVar(&flags.Seconds, "seconds", false, "Show seconds only")
	flag.BoolVar(&flags.Millis, "millis", false, "Show milliseconds only")
	flag.BoolVar(&flags.Nanos, "nanos", false, "Show nanoseconds only")
	flag.Parse()

	return flags
}
