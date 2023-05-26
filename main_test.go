package main

import "testing"

func Test_MemRead(t *testing.T) {
	c := newCpu()
	c.memory[0] = 1
	if c.memRead(0) != 1 {
		t.Errorf("should read 1")
	}
}

func Test_MemReadUint16(t *testing.T) {
	c := newCpu()
	c.memory[0] = 0xFF
	c.memory[1] = 0xFF
	if c.memReadUint16(0) != 0xFFFF {
		t.Errorf("should read 0xFFFF")
	}
}

func Test_MemWrite(t *testing.T) {
	c := newCpu()
	c.memWrite(0, 1)
	if c.memory[0] != 1 {
		t.Errorf("should write 1")
	}
}

func Test_MemWriteUint16(t *testing.T) {
	c := newCpu()
	c.memWriteUint16(0, 0xFFFF)
	if c.memory[0] != 0xFF || c.memory[1] != 0xFF {
		t.Errorf("should write 0xFFFF")
	}
}
