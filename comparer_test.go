package goroutree_test

import (
	"testing"

	"github.com/ScottMansfield/goroutree"
)

func TestInt(t *testing.T) {
	var a, b, c goroutree.Int = 4, 5, 6

	t.Run("Less", func(t *testing.T) {
		// Transitive
		if i, _ := a.Compare(b); i != -1 {
			t.Errorf("Expected %d to be less than %d and return -1", a, b)
		}

		if i, _ := b.Compare(c); i != -1 {
			t.Errorf("Expected %d to be less than %d and return -1", b, c)
		}

		if i, _ := a.Compare(c); i != -1 {
			t.Errorf("Expected %d to be less than %d and return -1", a, c)
		}
	})
	t.Run("Equal", func(t *testing.T) {
		if i, _ := a.Compare(a); i != 0 {
			t.Errorf("Expected %d to be less than %d and return 0", a, a)
		}

		if i, _ := b.Compare(b); i != 0 {
			t.Errorf("Expected %d to be less than %d and return 0", b, b)
		}

		if i, _ := c.Compare(c); i != 0 {
			t.Errorf("Expected %d to be less than %d and return 0", c, c)
		}
	})
	t.Run("Greater", func(t *testing.T) {
		// Transitive

		if i, _ := c.Compare(b); i != 1 {
			t.Errorf("Expected %d to be less than %d and return 1", c, b)
		}
		if i, _ := b.Compare(a); i != 1 {
			t.Errorf("Expected %d to be less than %d and return 1", b, a)
		}
		if i, _ := c.Compare(a); i != 1 {
			t.Errorf("Expected %d to be less than %d and return 1", c, a)
		}

	})
}

type NotInt string

func (NotInt) Compare(interface{}) (int, error) {
	return 0, nil
}

func TestIncompatable(t *testing.T) {
	var anInt goroutree.Int = 4
	var notAnInt NotInt = "Not An Int"

	if i, err := anInt.Compare(notAnInt); i != 0 {
		t.Errorf("Expected %d to not compare to %d and return 0 and \"%s\", got %d and \"%s\"", anInt, notAnInt, goroutree.NotComparable.Error(), i, err.Error())
	}
}
