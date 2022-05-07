package errorskit

import "fmt"

// Wrap is a drop-in replacement for errors.Wrap (https://github.com/pkg/errors) using STD's fmt.Errorf().
func Wrap(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}
