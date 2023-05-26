package main

type addrMode func(*cpu) uint16

func addrModeImplicit(c *cpu) uint16 {
	return 0
}

func addrModeAccumulator(c *cpu) uint16 {
	return 0
}

func addrModeImmediate(c *cpu) uint16 {
	addr := c.programCounter
	return addr
}

func addrModeZeroPage(c *cpu) uint16 {
	addr := c.memRead(c.programCounter)
	c.programCounter += 1
	return uint16(addr)
}

func addrModeZeroPageX(c *cpu) uint16 {
	addr := c.memRead(c.programCounter)
	c.programCounter += 1
	return uint16(wrappingSumUint8(addr, c.registerX))
}

func addrModeZeroPageY(c *cpu) uint16 {
	addr := c.memRead(c.programCounter)
	c.programCounter += 1
	return uint16(wrappingSumUint8(addr, c.registerY))
}

func addrModeRelative(c *cpu) uint16 {
	return 0
}

func addrModeAbsolute(c *cpu) uint16 {
	addr := c.memReadUint16(c.programCounter)
	c.programCounter += 2
	return addr
}

func addrModeAbsoluteX(c *cpu) uint16 {
	addr := c.memReadUint16(c.programCounter)
	c.programCounter += 2
	return uint16(wrappingSumUint16(addr, uint16(c.registerX)))
}

func addrModeAbsoluteY(c *cpu) uint16 {
	addr := c.memReadUint16(c.programCounter)
	c.programCounter += 2
	return uint16(wrappingSumUint16(addr, uint16(c.registerY)))
}

func addrModeIndirect(c *cpu) uint16 {
	return 0
}

func addrModeIndexedIndirect(c *cpu) uint16 {
	base := c.memRead(c.programCounter)
	c.programCounter += 1
	base = wrappingSumUint8(base, c.registerX)
	l := uint16(base)
	h := uint16(wrappingSumUint8(base, 1))
	return h<<8 | l
}

func addrModeIndirectIndexed(c *cpu) uint16 {
	base := c.memRead(c.programCounter)
	c.programCounter += 1
	base = wrappingSumUint8(base, c.registerX)
	l := uint16(base)
	h := uint16(wrappingSumUint8(base, 1))
	return wrappingSumUint16(h<<8|l, uint16(c.registerY))
}
