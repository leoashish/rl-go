// Package todo holds the one helper that every unfinished stub uses.
package todo

import (
	"errors"
	"fmt"
)

var ErrNotImplemented = errors.New("not implemented yet")

func NotImplemented(what string) error {
	return fmt.Errorf("%w: %s", ErrNotImplemented, what)
}
