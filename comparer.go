package goroutree

import "errors"

var NotComparable error = errors.New("Not Comparable")

type Comparer interface {
	// Compare yourself to something.  If the values are not comparable,
	// returns negative if the first argument is “less” than the second, zero
	// if they are “equal”, and positive if the first argument is “greater”
	// returns 0, and an error of "Not Comparable" if not comparable
	Compare(interface{}) (int, error)
}

type Int int

func (i Int) Compare(value interface{}) (int, error) {
	intValue, ok := value.(Int)
	if !ok {
		return 0, NotComparable
	}
	if i < intValue {
		return -1, nil
	} else if i == intValue {
		return 0, nil
	}
	return 1, nil
}
