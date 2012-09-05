// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Group is a group of nodes.
type Group struct {
	Parent *Group
	Nodes  []interface{}
}
