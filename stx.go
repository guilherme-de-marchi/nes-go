package main

const (
	STX_zeroPage  = 0x86
	STX_zeroPageY = 0x96
	STX_absolute  = 0x8E
)

type stx instruction

func (in stx) run() {
	addr := in.addrMode(in.cpu)
	in.cpu.memory[addr] = in.cpu.registerX
}

func (c *cpu) setSTX() {
	c.addInstruction(STX_zeroPage, stx(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(STX_zeroPageY, stx(newInstruction(c, addrModeZeroPageY)))
	c.addInstruction(STX_absolute, stx(newInstruction(c, addrModeAbsolute)))
}
