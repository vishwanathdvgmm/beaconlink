package version

import (
	"fmt"
	"runtime"
)

// These values are injected during release builds via ldflags.
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// Info contains BeaconLink build information.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
}

// Get returns the current build information.
func Get() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		BuildDate: Date,
		GoVersion: runtime.Version(),
	}
}

// String returns a human-readable representation.
func (i Info) String() string {
	return fmt.Sprintf(
		`BeaconLink

Version : %s
Commit  : %s
Build   : %s
Go      : %s`,
		i.Version,
		i.Commit,
		i.BuildDate,
		i.GoVersion,
	)
}
