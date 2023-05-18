package main

import "testing"

func Test_LDY_immediate(t *testing.T) {
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

func Test_LDY_zeroPage(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.loadAndRun([]uint8{
		LDY_zeroPage,
		0x01,
		BRK,
	})

	if c.registerY != 1 {
		t.Errorf("register Y has a wrong value: %b", c.registerY)
	}
}

func Test_LDY_zeroPageX(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.memory[0x02] = 2
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDY_zeroPageX,
		0x01,
		BRK,
	})

	if c.registerY != 2 {
		t.Errorf("register Y has a wrong value: %b", c.registerY)
	}
}

func Test_LDY_absolute(t *testing.T) {
	c := newCpu()
	c.memory[0x11FF] = 1
	c.loadAndRun([]uint8{
		LDY_absolute,
		0xFF,
		0x11,
		BRK,
	})

	if c.registerY != 1 {
		t.Errorf("register Y has a wrong value: %b", c.registerY)
	}
}

func Test_LDY_absoluteX(t *testing.T) {
	c := newCpu()
	c.memory[0x11F1] = 1
	c.memory[0x11F2] = 2
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDY_absoluteX,
		0xF1,
		0x11,
		BRK,
	})

	if c.registerY != 2 {
		t.Errorf("register Y has a wrong value: %b", c.registerY)
	}
}
