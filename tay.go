package main

const (
	TAY_implicit = 0xA8
)

type tay instruction

func (in tay) run() {
	in.cpu.registerY = in.cpu.registerA
	in.cpu.updateZNFlags(in.cpu.registerY)
}

func (c *cpu) setTAY() {
	c.addInstruction(TAY_implicit, tay(newInstruction(c, addrModeImplicit)))
}
