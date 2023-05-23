package main

const (
	LDX_immediate = 0xA2
	LDX_zeroPage  = 0xA6
	LDX_zeroPageY = 0xB6
	LDX_absolute  = 0xAE
	LDX_absoluteY = 0xBE
)

type ldx instruction

func (in ldx) run() {
	addr := in.addrMode(in.cpu)
	param := in.cpu.memRead(addr)
	in.cpu.programCounter += 1
	in.cpu.registerX = param
	in.cpu.updateZNFlags(in.cpu.registerX)
}

func (c *cpu) setLDX() {
	c.addInstruction(LDX_immediate, ldx(newInstruction(c, addrModeImmediate)))
	c.addInstruction(LDX_zeroPage, ldx(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(LDX_zeroPageY, ldx(newInstruction(c, addrModeZeroPageY)))
	c.addInstruction(LDX_absolute, ldx(newInstruction(c, addrModeAbsolute)))
	c.addInstruction(LDX_absoluteY, ldx(newInstruction(c, addrModeAbsoluteY)))
}
