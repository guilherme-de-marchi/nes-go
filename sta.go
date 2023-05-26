package main

const (
	STA_zeroPage        = 0x85
	STA_zeroPageX       = 0x95
	STA_absolute        = 0x8D
	STA_absoluteX       = 0x9D
	STA_absoluteY       = 0x99
	STA_indexedIndirect = 0x81
	STA_indirectIndexed = 0x91
)

type sta instruction

func (in sta) run() {
	addr := in.addrMode(in.cpu)
	in.cpu.memory[addr] = in.cpu.registerA
}

func (c *cpu) setSTA() {
	c.addInstruction(STA_zeroPage, sta(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(STA_zeroPageX, sta(newInstruction(c, addrModeZeroPageX)))
	c.addInstruction(STA_absolute, sta(newInstruction(c, addrModeAbsolute)))
	c.addInstruction(STA_absoluteX, sta(newInstruction(c, addrModeAbsoluteX)))
	c.addInstruction(STA_absoluteY, sta(newInstruction(c, addrModeAbsoluteY)))
	c.addInstruction(STA_indexedIndirect, sta(newInstruction(c, addrModeIndexedIndirect)))
	c.addInstruction(STA_indirectIndexed, sta(newInstruction(c, addrModeIndirectIndexed)))
}
