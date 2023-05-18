package main

import "testing"

func Test_LDX_immediate(t *testing.T) {
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

func Test_LDX_zeroPage(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.loadAndRun([]uint8{
		LDX_zeroPage,
		0x01,
		BRK,
	})

	if c.registerX != 1 {
		t.Errorf("register X has a wrong value: %b", c.registerX)
	}
}

func Test_LDX_zeroPageY(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.memory[0x02] = 2
	c.loadAndRun([]uint8{
		LDY_immediate,
		1,
		LDX_zeroPageY,
		0x01,
		BRK,
	})

	if c.registerX != 2 {
		t.Errorf("register X has a wrong value: %b", c.registerX)
	}
}

func Test_LDX_absolute(t *testing.T) {
	c := newCpu()
	c.memory[0x11FF] = 1
	c.loadAndRun([]uint8{
		LDX_absolute,
		0xFF,
		0x11,
		BRK,
	})

	if c.registerX != 1 {
		t.Errorf("register X has a wrong value: %b", c.registerX)
	}
}

func Test_LDX_absoluteY(t *testing.T) {
	c := newCpu()
	c.memory[0x11F1] = 1
	c.memory[0x11F2] = 2
	c.loadAndRun([]uint8{
		LDY_immediate,
		1,
		LDX_absoluteY,
		0xF1,
		0x11,
		BRK,
	})

	if c.registerX != 2 {
		t.Errorf("register X has a wrong value: %b", c.registerX)
	}
}
