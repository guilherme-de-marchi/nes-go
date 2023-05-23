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

// func Test_Load(t *testing.T) {
// 	c := newCpu()
// 	program := []uint8{
// 		LDA_immediate,
// 		0xFF,
// 		TAX,
// 		BRK,
// 	}
// 	c.load(program)

// 	s := c.memory[0x8000 : 0x8000+len(program)]
// 	for i, v := range program {
// 		if s[i] != v {
// 			t.Errorf("memory has wrong value on %v program line", i)
// 		}
// 	}
// }

// func Test_TAXWithData(t *testing.T) {
// 	c := newCpu()
// 	c.loadAndRun([]uint8{
// 		LDA_immediate,
// 		0xFF,
// 		TAX,
// 		BRK,
// 	})

// 	if c.registerA != 0xFF {
// 		t.Error("register X has a wrong value")
// 	}
// }

// func Test_INX(t *testing.T) {
// 	c := newCpu()
// 	c.loadAndRun([]uint8{
// 		LDA_immediate,
// 		0x00,
// 		TAX,
// 		INX,
// 		BRK,
// 	})

// 	if c.registerX != 1 {
// 		t.Error("should have added 1 to register X")
// 	}
// }

// func Test_INXOverflow(t *testing.T) {
// 	c := newCpu()
// 	c.loadAndRun([]uint8{
// 		LDA_immediate,
// 		0xFF,
// 		TAX,
// 		INX,
// 		INX,
// 		BRK,
// 	})

// 	if c.registerX != 1 {
// 		t.Error("should have added 1 to register X")
// 	}
// }

// func Test_UpdateZNFlagsWithData(t *testing.T) {
// 	c := newCpu()
// 	c.updateZNFlags(0xFF)

// 	if c.status&0b0000_0010 != 0 {
// 		t.Error("bit-1 should be 0")
// 	}

// 	if c.status&0b1000_0000 == 0 {
// 		t.Error("bit-7 should be 1")
// 	}
// }

// func Test_UpdateZNFlagsWithZero(t *testing.T) {
// 	c := newCpu()
// 	c.updateZNFlags(0x00)

// 	if c.status&0b0000_0010 == 0 {
// 		t.Error("bit-1 should be 1")
// 	}

// 	if c.status&0b1000_0000 != 0 {
// 		t.Error("bit-7 should be 0")
// 	}
// }

// func Test_WrappingSumUint8(t *testing.T) {
// 	if wrappingSumUint8(1, 1) != 2 {
// 		t.Error("should be 2")
// 	}

// 	if wrappingSumUint8(0xFF, 1) != 0 {
// 		t.Error("should be 0")
// 	}
// }

// func Test_WrappingSumUint16(t *testing.T) {
// 	if wrappingSumUint16(1, 1) != 2 {
// 		t.Error("should be 2")
// 	}

// 	if wrappingSumUint16(0xFFFF, 1) != 0 {
// 		t.Error("should be 0")
// 	}
// }
