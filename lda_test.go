package main

import "testing"

func TestLDA(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
		LDA_immediate,
		1,
		BRK,
	})

	if c.registerA != 1 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}
