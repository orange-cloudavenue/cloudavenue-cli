package customErrors

import "errors"

var (
	// Generic Error type
	ErrNotValidOutput  = errors.New("not valid output.")
	ErrNoHomeDirectory = errors.New("no home directory.")
	ErrNotImplemented  = errors.New("not implemented.")
	ErrConfigFile      = errors.New("error to configuration file.")
	ErrClient          = errors.New("error to initialize client.")
)

// IsNoHomeDirectory checks if the error is of type ErrNoHomeDirectory
func IsNoHomeDirectory(err error) bool {
	return errors.Is(err, ErrNoHomeDirectory)
}

// IsConfigFile checks if the error is of type ErrConfigFile
func IsConfigFile(err error) bool {
	return errors.Is(err, ErrConfigFile)
}

// IsClient checks if the error is of type ErrClient
func IsClient(err error) bool {
	return errors.Is(err, ErrClient)
}

// IsNotValidOutput checks if the error is of type ErrNotValidOutput
func IsNotValidOutput(err error) bool {
	return errors.Is(err, ErrNotValidOutput)
}
