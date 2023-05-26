package main

import "testing"

func TestLDX(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		BRK,
	})

	if c.registerX != 1 {
		t.Errorf("register X has a wrong value: %b", c.registerX)
	}
}
