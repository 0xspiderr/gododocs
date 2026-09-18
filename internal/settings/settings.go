package settings

// purpose: holds provided cmdline flags to provision
// other functions that might need them.
type Settings struct {
	OutDir string
	InDir  string
}

const (
	// default values for where the input is taken and
	// where the output will be stored is the current directory.
	DefaultInDir  = "."
	DefaultOutDir = "."
)
