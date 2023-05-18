package main

import "testing"

func Test_LDA_immediate(t *testing.T) {
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

func Test_LDA_zeroPage(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDA_zeroPage,
		0x01,
		BRK,
	})

	if c.registerA != 1 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}

func Test_LDA_zeroPageX(t *testing.T) {
	c := newCpu()
	c.memory[0x02] = 2
	c.registerX = 1
	c.resetProgramCounter()
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDA_zeroPageX,
		0x01,
		BRK,
	})

	if c.registerA != 2 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}

func Test_LDA_absolute(t *testing.T) {
	c := newCpu()
	c.memory[0x11FF] = 1
	c.loadAndRun([]uint8{
		LDA_absolute,
		0xFF,
		0x11,
		BRK,
	})

	if c.registerA != 1 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}

func Test_LDA_absoluteX(t *testing.T) {
	c := newCpu()
	c.memory[0x11FF] = 1
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDA_absoluteX,
		0xFE,
		0x11,
		BRK,
	})

	if c.registerA != 1 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}

// func Test_LDA_absoluteY(t *testing.T) {
// 	c := newCpu()
// 	c.memory[0x11FF] = 1
// 	c.loadAndRun([]uint8{
// 		LDY_immediate,
// 		1,
// 		LDA_absoluteY,
// 		0xFE,
// 		0x11,
// 		BRK,
// 	})

// 	if c.registerA != 1 {
// 		t.Errorf("register A has a wrong value: %b", c.registerA)
// 	}
// }

func Test_LDA_indexedIndirect(t *testing.T) {
	c := newCpu()
	c.memory[0x01] = 1
	c.memory[0x0302] = 2
	c.loadAndRun([]uint8{
		LDX_immediate,
		1,
		LDA_indexedIndirect,
		0x01,
		BRK,
	})

	if c.registerA != 2 {
		t.Errorf("register A has a wrong value: %b", c.registerA)
	}
}

// func Test_LDA_indirectIndexed(t *testing.T) {
// 	c := newCpu()
// 	c.memory[0x01] = 1
// 	c.memory[0x0302] = 2
// 	c.loadAndRun([]uint8{
// 		LDX_immediate,
// 		1,
// 		LDA_indexedIndirect,
// 		0x01,
// 		BRK,
// 	})

// 	if c.registerA != 2 {
// 		t.Errorf("register A has a wrong value: %b", c.registerA)
// 	}
// }
