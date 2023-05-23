package main

const (
	LDY_immediate = 0xA0
	LDY_zeroPage  = 0xA4
	LDY_zeroPageX = 0xB4
	LDY_absolute  = 0xAC
	LDY_absoluteX = 0xBC
)

type ldy instruction

func (in ldy) run() {
	addr := in.addrMode(in.cpu)
	param := in.cpu.memRead(addr)
	in.cpu.programCounter += 1
	in.cpu.registerY = param
	in.cpu.updateZNFlags(in.cpu.registerY)
}

func (c *cpu) setLDY() {
	c.addInstruction(LDY_immediate, ldy(newInstruction(c, addrModeImmediate)))
	c.addInstruction(LDY_zeroPage, ldy(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(LDY_zeroPageX, ldy(newInstruction(c, addrModeZeroPageX)))
	c.addInstruction(LDY_absolute, ldy(newInstruction(c, addrModeAbsolute)))
	c.addInstruction(LDY_absoluteX, ldy(newInstruction(c, addrModeAbsoluteX)))
}
