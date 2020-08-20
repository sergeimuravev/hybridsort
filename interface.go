package hybridsort

import (
	"sort"
)

// Interface is a collection type contract.
// Based on sort package interface.
type Interface interface {
	sort.Interface
	// ElementAt returns collection element by index.
	ElementAt(index int) interface{}
	// SetElementAt assigns element's value by index.
	SetElementAt(index int, elem interface{})
}
