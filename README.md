## Goroutree: A tree-based set made of coordinating goroutines

This project is an idea I had for a long time, but didn't have a reason to actually implement. It
was implemented for a blog post in the Gopher Academy Advent Series 2016.

(link to blog post to be added when it goes live)

It's a set that's made of a tree of coordinating goroutines. All communication is async with
responses returning via a channel supplied by the caller.

There's only three real operations: Insert, Contains, and Delete. 

* Insert will add an item if it does not already exist (and will tel you if it was successful).
* Contains will tell you whether the value is already in the tree.
* Delete will remove an item if it exists (and will tell you if it did).

There's a fourth, print, that is used to dump the state of the tree for verification and testing.

In general, each node in the tree is a separate goroutine that has a stream of messages coming in.
The node will respond to each message as it receives it and will forward to children as appropriate.
In order to start things out (and to deal with things like an empty set) there is a manager that
runs as a "super root" node and receives the messages first and forwards as necessary to the main
root node.

The code is fairly straightforward in terms of organization, just a single implementation file and a
single test file.

A note of caution: This is not production level code. It was written for a blog post as a toy and to
prove out the concept. I don't recommend running it in a production system without a lot more
testing in concurrent situations.
