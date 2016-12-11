//   Copyright 2016 Scott Mansfield
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package goroutree_test

import (
	"bytes"
	"testing"

	"github.com/ScottMansfield/goroutree"
)

func TestInsert(t *testing.T) {
	t.Run("One", func(t *testing.T) {
		g := goroutree.New()

		boolreschan := make(chan bool)

		g.Insert(boolreschan, 5)

		b := <-boolreschan
		if b != true {
			t.Fatalf("Expected a true result from inserting")
		}

		structreschan := make(chan struct{})
		buf := &bytes.Buffer{}

		g.Print(structreschan, buf)
		<-structreschan

		t.Logf("Printed tree: \n%s", buf.String())

		gold := "5\n"
		if buf.String() != gold {
			t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
		}
	})
	t.Run("Two", func(t *testing.T) {
		t.Run("SecondLess", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 5)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 4\n5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SecondGreater", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 5)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "5\n 6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SecondEqual", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 5)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != false {
				t.Fatalf("Expected a false result from inserting duplicate")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
	})
	t.Run("Three", func(t *testing.T) {
		t.Run("Balanced", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 5)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 4\n5\n 6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SkewLeft", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 6)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "  4\n 5\n6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SkewRight", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n 5\n  6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SkewRightLeft", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n  5\n 6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("SkewLeftRight", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 6)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 4\n  5\n6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("ThirdEqualRoot", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != false {
				t.Fatalf("Expected a false result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n 5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("ThirdEqualLeft", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 5)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 4)

			b = <-boolreschan
			if b != false {
				t.Fatalf("Expected a false result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 4\n5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("ThirdEqualRight", func(t *testing.T) {
			g := goroutree.New()

			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != false {
				t.Fatalf("Expected a false result from inserting")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n 5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
	})
	t.Run("BiggerTree", func(t *testing.T) {

		// Build up a bigger tree and test the printed output
		//
		//               7
		//       3                11
		//   1       5       9          13
		// 0   2   4   6   8   10    12    14

		//7,3,1,0,2,5,4,6,11,9,8,10,13,12,14
		g := goroutree.New()

		boolreschan := make(chan bool, 15)

		// create a balanced tree of 15 items
		g.Insert(boolreschan, 7)
		g.Insert(boolreschan, 3)
		g.Insert(boolreschan, 1)
		g.Insert(boolreschan, 0)
		g.Insert(boolreschan, 2)
		g.Insert(boolreschan, 5)
		g.Insert(boolreschan, 4)
		g.Insert(boolreschan, 6)
		g.Insert(boolreschan, 11)
		g.Insert(boolreschan, 9)
		g.Insert(boolreschan, 8)
		g.Insert(boolreschan, 10)
		g.Insert(boolreschan, 13)
		g.Insert(boolreschan, 12)
		g.Insert(boolreschan, 14)

		for i := 0; i < 15; i++ {
			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}
		}

		structreschan := make(chan struct{})
		buf := &bytes.Buffer{}

		g.Print(structreschan, buf)
		<-structreschan

		t.Logf("Printed tree: \n%s", buf.String())

		gold := "   0\n  1\n   2\n 3\n   4\n  5\n   6\n7\n   8\n  9\n   10\n 11\n   12\n  13\n   14\n"
		if buf.String() != gold {
			t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		g := goroutree.New()
		boolreschan := make(chan bool)
		g.Contains(boolreschan, 4)

		b := <-boolreschan
		if b != false {
			t.Fatal("Expected false result from contains")
		}
	})
	t.Run("One", func(t *testing.T) {
		t.Run("Hit", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Contains(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from contains")
			}
		})
		t.Run("Miss", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Contains(boolreschan, 5)

			b = <-boolreschan
			if b != false {
				t.Fatal("Expected false result from contains")
			}
		})
	})
	t.Run("Two", func(t *testing.T) {
		t.Run("Hit", func(t *testing.T) {
			t.Run("Root", func(t *testing.T) {
				g := goroutree.New()
				boolreschan := make(chan bool)

				g.Insert(boolreschan, 4)

				b := <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Insert(boolreschan, 6)

				b = <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Contains(boolreschan, 4)

				b = <-boolreschan
				if b != true {
					t.Fatal("Expected true result from contains")
				}
			})
			t.Run("Child", func(t *testing.T) {
				t.Run("Left", func(t *testing.T) {
					g := goroutree.New()
					boolreschan := make(chan bool)

					g.Insert(boolreschan, 4)

					b := <-boolreschan
					if b != true {
						t.Fatalf("Expected a true result from inserting")
					}

					g.Insert(boolreschan, 2)

					b = <-boolreschan
					if b != true {
						t.Fatalf("Expected a true result from inserting")
					}

					g.Contains(boolreschan, 2)

					b = <-boolreschan
					if b != true {
						t.Fatal("Expected true result from contains")
					}
				})
				t.Run("Right", func(t *testing.T) {
					g := goroutree.New()
					boolreschan := make(chan bool)

					g.Insert(boolreschan, 4)

					b := <-boolreschan
					if b != true {
						t.Fatalf("Expected a true result from inserting")
					}

					g.Insert(boolreschan, 6)

					b = <-boolreschan
					if b != true {
						t.Fatalf("Expected a true result from inserting")
					}

					g.Contains(boolreschan, 6)

					b = <-boolreschan
					if b != true {
						t.Fatal("Expected true result from contains")
					}
				})
			})
		})
		t.Run("Miss", func(t *testing.T) {
			t.Run("Left", func(t *testing.T) {
				g := goroutree.New()
				boolreschan := make(chan bool)

				g.Insert(boolreschan, 4)

				b := <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Contains(boolreschan, 3)

				b = <-boolreschan
				if b != false {
					t.Fatal("Expected true result from contains")
				}
			})
			t.Run("Right", func(t *testing.T) {
				g := goroutree.New()
				boolreschan := make(chan bool)

				g.Insert(boolreschan, 4)

				b := <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Contains(boolreschan, 5)

				b = <-boolreschan
				if b != false {
					t.Fatal("Expected true result from contains")
				}
			})
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		g := goroutree.New()
		boolreschan := make(chan bool)
		g.Delete(boolreschan, 4)

		b := <-boolreschan
		if b != false {
			t.Fatal("Expected false result from delete")
		}
	})
	t.Run("OneNode", func(t *testing.T) {
		t.Run("Hit", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("Miss", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 5)

			b = <-boolreschan
			if b != false {
				t.Fatal("Expected false result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
	})
	t.Run("OneChild", func(t *testing.T) {
		t.Run("LeftChild", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 3)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 3)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("RightChild", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := "4\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("Root", func(t *testing.T) {
			t.Run("PromoteLeft", func(t *testing.T) {

				g := goroutree.New()
				boolreschan := make(chan bool)

				g.Insert(boolreschan, 4)

				b := <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Insert(boolreschan, 3)

				b = <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Delete(boolreschan, 4)

				b = <-boolreschan
				if b != true {
					t.Fatal("Expected true result from delete")
				}

				structreschan := make(chan struct{})
				buf := &bytes.Buffer{}

				g.Print(structreschan, buf)
				<-structreschan

				t.Logf("Printed tree: \n%s", buf.String())

				gold := "3\n"
				if buf.String() != gold {
					t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
				}
			})
			t.Run("PromoteRight", func(t *testing.T) {

				g := goroutree.New()
				boolreschan := make(chan bool)

				g.Insert(boolreschan, 4)

				b := <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Insert(boolreschan, 5)

				b = <-boolreschan
				if b != true {
					t.Fatalf("Expected a true result from inserting")
				}

				g.Delete(boolreschan, 4)

				b = <-boolreschan
				if b != true {
					t.Fatal("Expected true result from delete")
				}

				structreschan := make(chan struct{})
				buf := &bytes.Buffer{}

				g.Print(structreschan, buf)
				<-structreschan

				t.Logf("Printed tree: \n%s", buf.String())

				gold := "5\n"
				if buf.String() != gold {
					t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
				}
			})
		})
	})
	t.Run("TwoChildren", func(t *testing.T) {
		t.Run("MinIsRightChild", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 3)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 3\n5\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("MinIsRightChildsLeftChild", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 3)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 3\n5\n 6\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
		t.Run("MinIsRightChildsLeftChildWithRightChild", func(t *testing.T) {
			g := goroutree.New()
			boolreschan := make(chan bool)

			g.Insert(boolreschan, 4)

			b := <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 3)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 7)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 5)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Insert(boolreschan, 6)

			b = <-boolreschan
			if b != true {
				t.Fatalf("Expected a true result from inserting")
			}

			g.Delete(boolreschan, 4)

			b = <-boolreschan
			if b != true {
				t.Fatal("Expected true result from delete")
			}

			structreschan := make(chan struct{})
			buf := &bytes.Buffer{}

			g.Print(structreschan, buf)
			<-structreschan

			t.Logf("Printed tree: \n%s", buf.String())

			gold := " 3\n5\n  6\n 7\n"
			if buf.String() != gold {
				t.Fatalf("Expected printed tree to be \"%s\" but got %s", gold, buf.String())
			}
		})
	})
}
