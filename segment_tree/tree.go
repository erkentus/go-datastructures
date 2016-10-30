package segment_tree

import (
	"errors"
	"fmt"
	"math"
)

const (
	//MaxPossibleNumber ...
	//Maximum possible value allowed to be stored in the segment tree
	MaxPossibleNumber = int(1e9 + 10)
)

/*
SegmentTree ...
    Data structure for quick range queries
*/
type SegmentTree struct {
	//tree values storing minimum values in the child nodes
	tree []int
	//lazy propagation values
	lazy []int
	//length of the original array
	origLen int
}

/*
New ...
   Initialize Segment Tree object by passing in the data array
   and building internal tree for future queries
*/
func New(x []int) (*SegmentTree, error) {
	if len(x) < 1 {
		return nil, errors.New("Provided slice should contain at least one element")
	}
	s := &SegmentTree{}
	n := 2*(1<<uint(math.Ceil(math.Log2(float64(len(x)))))) - 1
	s.tree = make([]int, n)
	s.lazy = make([]int, n)
	s.origLen = len(x)
	s.buildTree(x)
	return s, nil
}

/*
RangeMinQuery ...
    for a given range returns the minimum number existing in the original array
*/
func (s *SegmentTree) RangeMinQuery(l, r int) (int, error) {
	if l < 0 || l > r || r > s.origLen-1 {
		return 0, fmt.Errorf("Invalid range: %d to %d", l, r)
	}
	return s.recRMQ(0, l, r, 0, s.origLen-1), nil
}

/*
RangeAdd ...
    add x to all values in range [l,r]
*/
func (s *SegmentTree) RangeAdd(l, r int, x int) error {
	if l < 0 || l > r || r > s.origLen-1 {
		return fmt.Errorf("Invalid range: %d to %d", l, r)
	}
	s.recRangeAdd(0, l, r, 0, s.origLen-1, x)
	return nil
}

//
//Internal functions
//

func (s *SegmentTree) buildTree(x []int) {
	s.recBuild(0, 0, int(s.origLen-1), x)
}

func (s *SegmentTree) recBuild(cp, l, r int, x []int) {
	if l >= r {
		s.tree[cp] = x[l]
		return
	}
	mid := (l + r) / 2
	s.recBuild(2*cp+1, l, mid, x)
	s.recBuild(2*cp+2, mid+1, r, x)
	s.tree[cp] = min(s.tree[2*cp+1], s.tree[2*cp+2])
}

func (s *SegmentTree) recRMQ(cp, ql, qr, l, r int) int {
	if s.lazy[cp] != 0 {
		s.flush(cp, l, r)
	}
	if ql <= l && qr >= r {
		return s.tree[cp]
	}
	if ql > r || qr < l {
		return MaxPossibleNumber
	}
	mid := (l + r) / 2
	lmin := s.recRMQ(2*cp+1, ql, qr, l, mid)
	rmin := s.recRMQ(2*cp+2, ql, qr, mid+1, r)
	return min(lmin, rmin)
}

func (s *SegmentTree) recRangeAdd(cp, ql, qr, l, r, x int) {
	if s.lazy[cp] != 0 {
		s.flush(cp, l, r)
	}
	if ql > r || qr < l {
		return
	}
	if ql <= l && qr >= r {
		s.tree[cp] += x
		if l < r {
			s.lazy[2*cp+1] += x
			s.lazy[2*cp+2] += x
		}
		return
	}
	mid := (l + r) / 2
	s.recRangeAdd(2*cp+1, ql, qr, l, mid, x)
	s.recRangeAdd(2*cp+2, ql, qr, mid+1, r, x)
	s.tree[cp] = min(s.tree[2*cp+1], s.tree[2*cp+2])
}

func (s *SegmentTree) flush(cp, l, r int) {
	s.tree[cp] += s.lazy[cp]
	if l < r {
		s.lazy[2*cp+1] += s.lazy[cp]
		s.lazy[2*cp+2] += s.lazy[cp]
	}
	s.lazy[cp] = 0
}

//Helper functions
func min(args ...int) int {
	t := MaxPossibleNumber
	for _, x := range args {
		if x < t {
			t = x
		}
	}
	return t
}
