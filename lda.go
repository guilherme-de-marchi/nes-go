package main

const (
	LDA_immediate       = 0xA9
	LDA_zeroPage        = 0xA5
	LDA_zeroPageX       = 0xB5
	LDA_absolute        = 0xAD
	LDA_absoluteX       = 0xBD
	LDA_absoluteY       = 0xB9
	LDA_indexedIndirect = 0xA1
	LDA_indirectIndexed = 0xB1
)

type lda instruction

func (in lda) run() {
	addr := in.addrMode(in.cpu)
	param := in.cpu.memRead(addr)
	in.cpu.programCounter += 1
	in.cpu.registerA = param
	in.cpu.updateZNFlags(in.cpu.registerA)
}

func (c *cpu) setLDA() {
	c.addInstruction(LDA_immediate, lda(newInstruction(c, addrModeImmediate)))
	c.addInstruction(LDA_zeroPage, lda(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(LDA_zeroPageX, lda(newInstruction(c, addrModeZeroPageX)))
	c.addInstruction(LDA_absolute, lda(newInstruction(c, addrModeAbsolute)))
	c.addInstruction(LDA_absoluteX, lda(newInstruction(c, addrModeAbsoluteX)))
	c.addInstruction(LDA_absoluteY, lda(newInstruction(c, addrModeAbsoluteY)))
	c.addInstruction(LDA_indexedIndirect, lda(newInstruction(c, addrModeIndexedIndirect)))
	c.addInstruction(LDA_indirectIndexed, lda(newInstruction(c, addrModeIndirectIndexed)))
}
