package main

import "testing"

func TestSTA(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
		LDA_immediate,
		1,
		STA_zeroPage,
		0x01,
		BRK,
	})

	if c.memory[0x01] != 1 {
		t.Errorf("value in memory is not 1")
	}
}
