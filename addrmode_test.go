package main

import "testing"

func TestImmediate(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = programAddr

	c := newCpu()
	c.load([]uint8{exp})
	c.resetProgramCounter()

	addr := addrModeImmediate(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestZeroPage(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x01

	c := newCpu()
	c.memory[target] = exp
	c.load([]uint8{uint8(target)})
	c.resetProgramCounter()

	addr := addrModeZeroPage(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestZeroPageX(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x02

	c := newCpu()
	c.memory[target] = exp
	c.registerX = 1
	c.load([]uint8{0x01})
	c.resetProgramCounter()

	addr := addrModeZeroPageX(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestZeroPageY(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x02

	c := newCpu()
	c.memory[target] = exp
	c.registerY = 1
	c.load([]uint8{0x01})
	c.resetProgramCounter()

	addr := addrModeZeroPageY(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestAbsolute(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x0201

	c := newCpu()
	c.memory[target] = exp
	c.load([]uint8{
		0x01,
		0x02,
	})
	c.resetProgramCounter()

	addr := addrModeAbsolute(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestAbsoluteX(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x0202

	c := newCpu()
	c.memory[target] = exp
	c.registerX = 1
	c.load([]uint8{
		0x01,
		0x02,
	})
	c.resetProgramCounter()

	addr := addrModeAbsoluteX(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestAbsoluteY(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x0202

	c := newCpu()
	c.memory[target] = exp
	c.registerY = 1
	c.load([]uint8{
		0x01,
		0x02,
	})
	c.resetProgramCounter()

	addr := addrModeAbsoluteY(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestIndexedIndirect(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x0302

	c := newCpu()
	c.memory[target] = exp
	c.registerX = 1
	c.load([]uint8{0x01})
	c.resetProgramCounter()

	addr := addrModeIndexedIndirect(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}

func TestIndirectIndexed(t *testing.T) {
	var exp uint8 = 1
	var target uint16 = 0x0303

	c := newCpu()
	c.memory[target] = exp
	c.registerX = 1
	c.registerY = 1
	c.load([]uint8{0x01})
	c.resetProgramCounter()

	addr := addrModeIndirectIndexed(c)
	if addr != target {
		t.Errorf("unexpected addr %X", addr)
	}

	value := c.memRead(addr)
	if value != exp {
		t.Errorf("unexpected value %X at addr %X", value, addr)
	}
}
