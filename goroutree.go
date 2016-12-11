package goroutree

import (
	"bytes"
	"fmt"
	"io"
)

///////////////////////
// Command definitions
///////////////////////

type cmdType int

const (
	ctInvalid cmdType = iota
	ctInsert
	ctContains
	ctDelete
	ctPrint
	ctNewChild
	ctExtractMin
)

func (ct cmdType) String() string {
	switch ct {
	case ctInvalid:
		return "ctInvalid"
	case ctInsert:
		return "ctInsert"
	case ctContains:
		return "ctContains"
	case ctDelete:
		return "ctDelete"
	case ctPrint:
		return "ctPrint"
	case ctNewChild:
		return "ctNewChild"
	case ctExtractMin:
		return "ctExtractMin"
	default:
		panic("unrecognized command type")
	}
}

type cmd interface {
	typ() cmdType
}

type insertCmd struct {
	reschan chan bool
	val     int
}

func (c insertCmd) typ() cmdType {
	return ctInsert
}

type containsCmd struct {
	reschan chan bool
	val     int
}

func (c containsCmd) typ() cmdType {
	return ctContains
}

type deleteCmd struct {
	reschan chan bool
	val     int
	left    bool
}

func (c deleteCmd) typ() cmdType {
	return ctDelete
}

type printCmd struct {
	reschan chan struct{}
	level   int
	w       io.Writer
}

func (c printCmd) typ() cmdType {
	return ctPrint
}

type newChildCmd struct {
	childchan chan cmd
	left      bool
}

func (c newChildCmd) typ() cmdType {
	return ctNewChild
}

type extractMinCmd struct {
	reschan chan subtreeMinResponse
	first   bool
}

func (c extractMinCmd) typ() cmdType {
	return ctExtractMin
}

type subtreeMinResponse struct {
	val       int
	newchild  bool
	childchan chan cmd
}

////////////////////////////
// Goroutree implementation
////////////////////////////

// Goroutree is a tree set represented by a set of running goroutines, one per
// node in the tree. It is essentially an actor-based tree that holds
// machine-sized integers. The tree is unbalanced and does no rotations, so a
// series of inserts and deletes can make it very unbalanced. The tree is not
// going to be bombproof (because this is for a blog post) and probably has some
// obvious races.
type Goroutree struct {
	cmdchan chan cmd
}

// New creates a new empty Goroutree
func New() *Goroutree {
	cmdchan := make(chan cmd)
	go manager(cmdchan)

	return &Goroutree{cmdchan}
}

func manager(main chan cmd) {
	var cmdchan chan cmd

	for c := range main {
		switch c.typ() {
		case ctInsert:
			if cmdchan == nil {
				ic := c.(insertCmd)
				cmdchan = spawn(ic.val, main)
				ic.reschan <- true
				continue
			}

			cmdchan <- c

		case ctContains:
			if cmdchan == nil {
				cc := c.(containsCmd)
				cc.reschan <- false
				continue
			}

			cmdchan <- c

		case ctDelete:
			if cmdchan == nil {
				dc := c.(deleteCmd)
				dc.reschan <- false
				continue
			}

			cmdchan <- c

		case ctPrint:
			if cmdchan == nil {
				pc := c.(printCmd)
				pc.w.Write([]byte("\n"))
				pc.reschan <- struct{}{}
			}

			cmdchan <- c

		case ctNewChild:
			nc := c.(newChildCmd)
			cmdchan = nc.childchan

		default:
			panic(fmt.Sprintf("UNEXPECTED COMMAND: %#v", c))
		}
	}
}

// Insert adds a new value into the set if it does not already exist. The channel
// passed will receive a true if the value was successfully inserted and a false
// if the value already existed.
func (g *Goroutree) Insert(reschan chan bool, val int) {
	g.cmdchan <- insertCmd{
		reschan: reschan,
		val:     val,
	}
}

// Contains will tell if the set contains the given value. The channel passed will
// receive a true if the value does exist in the set and a false if not.
func (g *Goroutree) Contains(reschan chan bool, val int) {
	g.cmdchan <- containsCmd{
		reschan: reschan,
		val:     val,
	}
}

// Delete removes a value from the tree set if it exists. The channel passed will
// receive a true if the value did exist in the set and a false if not.
func (g *Goroutree) Delete(reschan chan bool, val int) {
	g.cmdchan <- deleteCmd{
		reschan: reschan,
		val:     val,
	}
}

// Print will print out the tree. This is a blocking operation, so no other
// messages can be processed while printing. This is for debugging purposes only.
func (g *Goroutree) Print(reschan chan struct{}, w io.Writer) {
	g.cmdchan <- printCmd{
		reschan: reschan,
		w:       w,
	}
}

// spawn creates a new node that owns a value.
// within this function is the logic that each node runs. Essentially it is an
// infinite loop that responds to messages sent on its command channel. It then
// decides to either act on that message or pass it on down the tree.
func spawn(val int, parentchan chan cmd) chan cmd {

	cmdchan := make(chan cmd)

	// cmdchan will be a stream of commands to be done in this node
	// val is a constant value that this node holds
	go func(cmdchan, parentchan chan cmd, val int) {
		var left, right chan cmd

		for cm := range cmdchan {
			switch cm.typ() {
			case ctInsert:
				c := cm.(insertCmd)

				if c.val == val {
					c.reschan <- false
					continue
				}

				// left branch
				if c.val < val {
					// if the left node exists, send it down.
					if left != nil {
						left <- c
						continue
					}

					left = spawn(c.val, cmdchan)
					c.reschan <- true
					continue
				}

				// right branch
				if right != nil {
					right <- c
					continue
				}

				right = spawn(c.val, cmdchan)
				c.reschan <- true

			case ctContains:
				c := cm.(containsCmd)

				if c.val == val {
					c.reschan <- true
					continue
				}

				// Go right if the value is bigger,
				// left if smaller
				if c.val > val && right != nil {
					right <- c
					continue
				}

				if left != nil {
					left <- c
					continue
				}

				// if we get here, the value does not exist in the tree
				c.reschan <- false

			case ctDelete:
				c := cm.(deleteCmd)

				// if a match, delete this node.
				if c.val == val {

					// if this is a leaf node with no children, it just returns
					if left == nil && right == nil {
						// send death message to parent
						parentchan <- newChildCmd{
							left:      c.left,
							childchan: nil,
						}

						c.reschan <- true
						return
					}

					// one child, promote it to current position by sending parent a message
					// we know at this point that one is not nil, so this checks if we have
					// one and only one not nil child.
					if left == nil || right == nil {

						var childchan chan cmd
						if left != nil {
							childchan = left
						} else {
							childchan = right
						}

						// promote child
						parentchan <- newChildCmd{
							left:      c.left,
							childchan: childchan,
						}

						c.reschan <- true
						return
					}

					// At this point, we need to substitute the current node with either the
					// maximum node on the left subtree or the minimum node on the right subtree.
					// For simplicity, this implementation always chooses to pull the minimum node
					// out of the right subtree.
					// The pattern is to send a message down the right subtree to find the minimum
					// node. Once it's found, it will have either one child or none. In this special
					// case, the node that is found will take care of removing itself and send its
					// value back to this node. TO make things simpler, this node will simply take
					// the value and assign it as its owned value. I could do some trickery with
					// reassigning channels all over the place to physically transplant that other
					// node to this position, but that just seems silly to do if I can get away with
					// just taking ownership of that value.

					reschan := make(chan subtreeMinResponse)
					right <- extractMinCmd{
						reschan: reschan,
						first:   true,
					}

					res := <-reschan
					val = res.val

					if res.newchild {
						right = res.childchan
					}

					c.reschan <- true
					continue
				}

				if c.val > val && right != nil {
					c.left = false
					right <- c
					continue
				}

				if left != nil {
					c.left = true
					left <- c
					continue
				}

				// if we get here, the value does not exist in the tree
				c.reschan <- false

			case ctExtractMin:
				c := cm.(extractMinCmd)

				// The first node might be the one that we want, in which case
				// we need to know. That node will be the right child and not
				// the left like all others
				if left != nil {
					if c.first {
						c.first = false
					}
					left <- c
					continue
				}

				// this is the right child of the node that is being deleted
				// it needs special attention here because there can be a race
				// between the parentchan message below and an external command.
				// The same place we're trying to send the parentchan message is
				// waiting on the reschan, so this special message takes care of
				// both at once.
				if c.first {
					c.reschan <- subtreeMinResponse{
						val:       val,
						newchild:  true,
						childchan: right,
					}
				}

				// then replace self at parent with whatever is at the right
				// nil is fine here, so no check
				parentchan <- newChildCmd{
					left:      !c.first,
					childchan: right,
				}

				// send back the min value for this subtree
				c.reschan <- subtreeMinResponse{val: val}

				return

			case ctNewChild:
				c := cm.(newChildCmd)

				println(c.left, c.childchan)

				if c.left {
					left = c.childchan
					continue
				}

				right = c.childchan

			case ctPrint:
				// Inorder printing traversal of the tree
				c := cm.(printCmd)

				// make a new command for the children. FOR THE CHILDREN.
				// Each child gets the command and a chance to finish its work before
				// this node mvoes on. This means that printing the tree is basically
				// a blocking operation in which no other operations can be done.
				childcmd := c
				childcmd.reschan = make(chan struct{})
				childcmd.level++

				if left != nil {
					left <- childcmd
					<-childcmd.reschan
				}

				// no this is not very efficient, but this is for debugging
				indent := bytes.Repeat([]byte(" "), c.level)
				fmt.Fprintf(c.w, "%s%d\n", indent, val)

				if right != nil {
					right <- childcmd
					<-childcmd.reschan
				}

				c.reschan <- struct{}{}

			default:
				panic(fmt.Sprintf("UNEXPECTED COMMAND: %#v", cm))
			}
		}
	}(cmdchan, parentchan, val)

	return cmdchan
}
