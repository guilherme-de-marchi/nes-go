package main

const (
	STY_zeroPage  = 0x84
	STY_zeroPageX = 0x94
	STY_absolute  = 0x8C
)

type sty instruction

func (in sty) run() {
	addr := in.addrMode(in.cpu)
	in.cpu.memory[addr] = in.cpu.registerY
}

func (c *cpu) setSTY() {
	c.addInstruction(STY_zeroPage, stx(newInstruction(c, addrModeZeroPage)))
	c.addInstruction(STY_zeroPageX, stx(newInstruction(c, addrModeZeroPageX)))
	c.addInstruction(STY_absolute, stx(newInstruction(c, addrModeAbsolute)))
}
