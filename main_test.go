package main

import "testing"

func TestLoad(t *testing.T) {
	c := newCpu()
	program := []uint8{
		LDA,
		0xFF,
		TAX,
		BRK,
	}
	c.load(program)

	s := c.memory[0x8000 : 0x8000+len(program)]
	for i, v := range program {
		if s[i] != v {
			t.Errorf("memory has wrong value on %v program line", i)
		}
	}
}

func TestTAXWithData(t *testing.T) {
	c := newCpu()
	c.loadAndRun([]uint8{
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
	c.loadAndRun([]uint8{
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
	c.loadAndRun([]uint8{
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
	c.loadAndRun([]uint8{
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
