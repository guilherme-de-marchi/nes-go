package main

import "testing"

func TestLDY(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
		LDY_immediate,
		1,
		BRK,
	})

	if c.registerY != 1 {
		t.Errorf("register Y has a wrong value: %b", c.registerY)
	}
}
