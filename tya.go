package main

const (
	TYA_implicit = 0xAA
)

type tya instruction

func (in tya) run() {
	in.cpu.registerA = in.cpu.registerY
	in.cpu.updateZNFlags(in.cpu.registerA)
}

func (c *cpu) setTYA() {
	c.addInstruction(TYA_implicit, tya(newInstruction(c, addrModeImplicit)))
}
