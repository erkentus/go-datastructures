package segment_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type query struct {
	l        int
	r        int
	res      int
	isupdate bool
}

type testpair struct {
	vals    []int
	queries []query
}

// normal behaviour
var tp1 = testpair{
	vals:    []int{1, 2, 3, 4, 5, 6},
	queries: []query{{0, 0, 1, false}, {1, 1, 2, false}, {0, 3, 1, false}, {3, 5, 4, false}},
}

// empty array
var tp2 = testpair{
	vals:    []int{},
	queries: []query{},
}

// invalid range
var tp3 = testpair{
	vals:    []int{-100, 0, 100},
	queries: []query{{0, 3, 1e9, false}, {2, 1, 1e9, false}},
}

// range updates + check rmq result
var tp4 = testpair{
	vals:    []int{1, 2, 3, 4, 5, 6, 7, 8},
	queries: []query{{0, 0, 1, true}, {0, 1, 2, false}, {0, 3, 1, true}, {0, 4, 3, false}, {3, 5, 5, false}},
}

// invalid range updates
var tp5 = testpair{
	vals:    []int{1, 2, 3, 4, 5},
	queries: []query{{0, -1, 1, true}, {0, 10, 10, true}},
}

func TestRangeMinQuery(t *testing.T) {
	st1, _ := New(tp1.vals)
	for _, rq := range tp1.queries {
		v, err := st1.RangeMinQuery(rq.l, rq.r)
		assert.Equal(t, v, rq.res)
		assert.Equal(t, err, nil)
	}

	_, err := New(tp2.vals)
	assert.EqualError(t, err, "Provided slice should contain at least one element")

	st3, err := New(tp3.vals)
	for _, rq := range tp3.queries {
		v, err := st3.RangeMinQuery(rq.l, rq.r)
		assert.EqualError(t, err, fmt.Sprintf("Invalid range: %d to %d", rq.l, rq.r))
		assert.Equal(t, v, 0)
	}
}

func TestRangeAdd(t *testing.T) {
	st4, _ := New(tp4.vals)
	for _, uq := range tp4.queries {
		if uq.isupdate {
			err := st4.RangeAdd(uq.l, uq.r, uq.res)
			assert.Equal(t, err, nil)
		} else {
			v, err := st4.RangeMinQuery(uq.l, uq.r)
			assert.Equal(t, err, nil)
			assert.Equal(t, v, uq.res)
		}
	}
	st5, _ := New(tp5.vals)
	for _, uq := range tp5.queries {
		if uq.isupdate {
			err := st5.RangeAdd(uq.l, uq.r, uq.res)
			assert.EqualError(t, err, fmt.Sprintf("Invalid range: %d to %d", uq.l, uq.r))
		}
	}
}
