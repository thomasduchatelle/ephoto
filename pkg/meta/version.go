// Package meta expose information about the application like the version.
package meta

const (
	version = "1.3.0"
)

// Version returns the version of the app, updated by ci/pre-release.sh
func Version() string {
	return version
}
