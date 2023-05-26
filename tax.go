package main

const (
	TAX_implicit = 0xAA
)

type tax instruction

func (in tax) run() {
	in.cpu.registerX = in.cpu.registerA
	in.cpu.updateZNFlags(in.cpu.registerX)
}

func (c *cpu) setTAX() {
	c.addInstruction(TAX_implicit, tax(newInstruction(c, addrModeImplicit)))
}
