package main

const (
	TXA_implicit = 0xAA
)

type txa instruction

func (in txa) run() {
	in.cpu.registerA = in.cpu.registerX
	in.cpu.updateZNFlags(in.cpu.registerA)
}

func (c *cpu) setTXA() {
	c.addInstruction(TXA_implicit, txa(newInstruction(c, addrModeImplicit)))
}
