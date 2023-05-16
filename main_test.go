package main

import "testing"

func TestTAXWithData(t *testing.T) {
	c := newCpu()
	c.interpret([]uint8{
		LDA,
		0xFF,
		TAX,
		BRK,
	})

	if c.registerA != 0xFF {
		t.Error("register X has a wrong value")
	}
}

func TestLDA(t *testing.T) {
	c := newCpu()
	c.interpret([]uint8{
		LDA,
		0xFF,
		BRK,
	})

	if c.registerA != 0xFF {
		t.Error("register A has a wrong value")
	}
}

func TestINX(t *testing.T) {
	c := newCpu()
	c.interpret([]uint8{
		LDA,
		0x00,
		TAX,
		INX,
		BRK,
	})

	if c.registerX != 1 {
		t.Error("should have added 1 to register X")
	}
}

func TestINXOverflow(t *testing.T) {
	c := newCpu()
	c.interpret([]uint8{
		LDA,
		0xFF,
		TAX,
		INX,
		INX,
		BRK,
	})

	if c.registerX != 1 {
		t.Error("should have added 1 to register X")
	}
}

func TestUpdateZNFlagsWithData(t *testing.T) {
	c := newCpu()
	c.updateZNFlags(0xFF)

	if c.status&0b0000_0010 != 0 {
		t.Error("bit-1 should be 0")
	}

	if c.status&0b1000_0000 == 0 {
		t.Error("bit-7 should be 1")
	}
}

func TestUpdateZNFlagsWithZero(t *testing.T) {
	c := newCpu()
	c.updateZNFlags(0x00)

	if c.status&0b0000_0010 == 0 {
		t.Error("bit-1 should be 1")
	}

	if c.status&0b1000_0000 != 0 {
		t.Error("bit-7 should be 0")
	}
}
