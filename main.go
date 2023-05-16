package main

const (
	BRK = 0x00
	LDA = 0xA9
	TAX = 0xAA
	INX = 0xE8
)

type cpu struct {
	registerA,
	registerX,
	status uint8
	programCounter uint16
	memory         [0xFFFF]uint8
}

func newCpu() *cpu {
	return &cpu{}
}

func (c *cpu) memRead(addr uint16) uint8 {
	return c.memory[addr]
}

func (c *cpu) memWrite(addr uint16, data uint8) {
	c.memory[addr] = data
}

func (c *cpu) load(program []uint8) {
	for i, v := range program {
		c.memory[0x8000+i] = v
	}
	c.programCounter = 0x8000
}

func (c *cpu) run() {
	for {
		opsCode := c.memRead(c.programCounter)
		c.programCounter += 1

		switch opsCode {
		case BRK:
			return

		case TAX:
			c.TAX()
			c.updateZNFlags(c.registerX)

		case LDA:
			param := c.memRead(c.programCounter)
			c.programCounter += 1
			c.LDA(param)
			c.updateZNFlags(c.registerA)

		case INX:
			c.INX()
			c.updateZNFlags(c.registerX)
		}
	}
}

func (c *cpu) loadAndRun(program []uint8) {
	c.load(program)
	c.run()
}

func (c *cpu) TAX() {
	c.registerX = c.registerA
}

func (c *cpu) LDA(param uint8) {
	c.registerA = param
}

func (c *cpu) INX() {
	c.registerX += 1
}

func (c *cpu) updateZNFlags(result uint8) {
	if result == 0 {
		c.status = c.status | 0b0000_0010
	} else {
		c.status = c.status & 0b1111_1101
	}

	if result&0b1000_0000 != 0 {
		c.status = c.status | 0b1000_0000
	} else {
		c.status = c.status & 0b0111_1111
	}
}

func main() {
}
