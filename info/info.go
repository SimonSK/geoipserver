package info

import "fmt"

const (
	Name         = "MaxMind GeoIP2 WebAPI Server"
	Description  = "A simple web API server that reads GeoIP2/GeoLite2 database binary and echoes out latitude and longitude information for a given IP address"
	Author       = "Simon Kim"
	VersionMajor = 0
	VersionMinor = 1
)

var (
	Version         = fmt.Sprintf("%d.%d", VersionMajor, VersionMinor)
	NameWithVersion = fmt.Sprintf("%s v%s", Name, Version)
)
