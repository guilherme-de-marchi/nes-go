package main

import "testing"

func TestSTX(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		STX_zeroPage,
		0x01,
		BRK,
	})

	if c.memory[0x01] != 1 {
		t.Errorf("value in memory is not 1")
	}
}
