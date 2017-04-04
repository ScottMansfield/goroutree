package goroutree_test

import (
	"testing"

	"github.com/ScottMansfield/goroutree"
)

func TestConcurrent(t *testing.T) {
	g := goroutree.New()
	boolreschan := make(chan bool)
	var vals map[goroutree.Int]bool = make(map[goroutree.Int]bool)

	for i := -100000; i <= 100000; i++ {
		vals[goroutree.Int(i)] = true
	}

	for k, _ := range vals {
		g.Insert(boolreschan, k)
	}

	for k, _ := range vals {
		g.Delete(boolreschan, k)
	}
}
