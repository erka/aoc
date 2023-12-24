package depth

import (
	"gonum.org/v1/gonum/graph"
)

// Int64s is a set of int64 identifiers.
type Int64s map[int64]struct{}

// The simple accessor methods for Ints are provided to allow ease of
// implementation change should the need arise.
var intDefaultValue = struct{}{}

// Add inserts an element into the set.
func (s Int64s) Add(e int64) {
	s[e] = intDefaultValue
}

// Has reports the existence of the element in the set.
func (s Int64s) Has(e int64) bool {
	_, ok := s[e]
	return ok
}

// Remove deletes the specified element from the set.
func (s Int64s) Remove(e int64) {
	delete(s, e)
}

// Count reports the number of elements stored in the set.
func (s Int64s) Count() int {
	return len(s)
}

// DepthLast implements stateful depth-last graph traversal.
type DepthLast struct {
	// Visit is called on all nodes on their first visit.
	Visit func(graph.Node)

	// Traverse is called on all edges that may be traversed
	// during the walk. This includes edges that would hop to
	// an already visited node.
	//
	// The value returned by Traverse determines whether an
	// edge can be traversed during the walk.
	Traverse func(graph.Edge) bool

	// stack   NodeStack
	visited Int64s
}

// Walk performs a depth traversal of the graph g starting from the given node,
func (d *DepthLast) WalkAll(g graph.Graph, from graph.Node, end graph.Node, depth func(int)) {
	if d.visited == nil {
		d.visited = make(Int64s)
	}
	u := from
	uid := u.ID()
	if d.visited.Has(uid) {
		return
	}
	if d.Visit != nil {
		d.Visit(u)
	}
	if uid == end.ID() {
		depth(d.visited.Count())
	}
	d.visited.Add(uid)
	to := g.From(uid)
	for to.Next() {
		v := to.Node()
		vid := v.ID()
		if d.Traverse != nil && !d.Traverse(g.Edge(uid, vid)) {
			continue
		}
		d.WalkAll(g, v, end, depth)
	}
	d.visited.Remove(uid)
}
