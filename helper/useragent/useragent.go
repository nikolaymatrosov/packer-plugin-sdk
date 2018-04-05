package useragent

import (
	"fmt"
	"runtime"

	"github.com/hashicorp/packer/version"
)

var (
	// projectURL is the project URL.
	projectURL = "https://www.packer.io/"

	// rt is the runtime - variable for tests.
	rt = runtime.Version()

	// versionFunc is the func that returns the current version. This is a
	// function to take into account the different build processes and distinguish
	// between enterprise and oss builds.
	versionFunc = func() string {
		return version.FormattedVersion()
	}
)

// String returns the consistent user-agent string for Packer.
func String() string {
	return fmt.Sprintf("Packer/%s (+%s; %s)",
		versionFunc(), projectURL, rt)
}
