package errs

import (
	"errors"
	"fmt"
)

var (
	OutOfBounds = errors.New("out of bounds")
	HeightZero  = errors.New("unallowed on height 0")
	Duplicated  = errors.New("value is already present in the tree")
)

func OpOnZeroLen(operation, dataStructure string) error {
	err := fmt.Sprintf("can't %s  %s of lenght 0", operation, dataStructure)
	return errors.New(err)
}
