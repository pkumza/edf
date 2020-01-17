package edf

import (
	"testing"
)

func Test_A(t *testing.T) {
	// Some items and their priorities.
	A := &Entry{Value: "✅ ", Weight: 120}
	B := &Entry{Value: "❎ ", Weight: 120}
	C := &Entry{Value: "✅ ", Weight: 120}
	D := &Entry{Value: "⭕️ ", Weight: 11}
	E := &Entry{Value: "🚷 ", Weight: 11}
	F := &Entry{Value: "🔞 ", Weight: 11}

	for loop := 0; loop < 10; loop++ {
		// Create a priority queue, put the items in it, and
		// establish the priority queue (heap) invariants.
		e := NewEDF([]*Entry{A, B, C, D, E, F})

		// Take the items out; they arrive in increasing deadline order.
		seq := ""
		for i := 0; i < 60; i++ {
			item := e.Pick()
			seq += item.Value
		}
		t.Logf("%s ", seq)
	}
}
