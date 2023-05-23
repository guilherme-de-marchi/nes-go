package main

import (
	"log"
)

type instructionFunc func(*cpu, addrMode)

type instruction struct {
	cpu      *cpu
	addrMode addrMode
}

func newInstruction(c *cpu, mode addrMode) instruction {
	return instruction{
		cpu:      c,
		addrMode: mode,
	}
}

type instructioner interface {
	run()
}

const (
// lda = iota
// ldx
)

const (
	BRK = 0x00

	STA_zeroPage        = 0x85
	STA_zeroPageX       = 0x95
	STA_absolute        = 0x8D
	STA_absoluteX       = 0x9D
	STA_absoluteY       = 0x99
	STA_indexedIndirect = 0x81
	STA_indirectIndexed = 0x91

	STX_zeroPage  = 0x86
	STX_zeroPageY = 0x96
	STX_absolute  = 0x8E

	TAX = 0xAA
	INX = 0xE8

	memorySize     = 0xFFFF
	programAddr    = 0x8000
	programRefAddr = 0xFFFC
)

type cpu struct {
	registerA,
	registerX,
	registerY,
	status uint8
	programCounter uint16
	memory         [memorySize]uint8
	instructions   map[uint8]instructioner
}

func newCpu() *cpu {
	return &cpu{
		instructions: make(map[uint8]instructioner),
	}
}

func (c *cpu) addInstruction(opCode uint8, in instructioner) {
	c.instructions[opCode] = in
}

func (c *cpu) memRead(addr uint16) uint8 {
	return c.memory[addr]
}

func (c *cpu) memWrite(addr uint16, data uint8) {
	c.memory[addr] = data
}

func (c *cpu) memReadUint16(addr uint16) uint16 {
	l := uint16(c.memory[addr])
	h := uint16(c.memory[addr+1])
	return h<<8 | l
}

func (c *cpu) memWriteUint16(addr, data uint16) {
	l := uint8(data & 0x00FF)
	h := uint8(data >> 8)
	c.memory[addr] = l
	c.memory[addr+1] = h
}

func (c *cpu) reset() {
	c.registerA = 0
	c.registerX = 0
	c.registerY = 0
	c.status = 0
	c.resetProgramCounter()
	c.setLDA()
	c.setLDX()
	c.setLDY()
}

func (c *cpu) resetProgramCounter() {
	c.programCounter = c.memReadUint16(programRefAddr)
}

func (c *cpu) load(program []uint8) {
	for i, v := range program {
		c.memory[programAddr+i] = v
	}
	c.memWriteUint16(programRefAddr, programAddr)
}

func (c *cpu) run() {
	for {
		op := c.memRead(c.programCounter)
		c.programCounter += 1

		if op == BRK {
			break
		}

		in, ok := c.instructions[op]
		if !ok {
			log.Fatal("no instruction found: ", op)
			return
		}

		in.run()

		// switch op {
		// case BRK:
		// 	return

		// case TAX:
		// 	c.TAX()
		// 	c.updateZNFlags(c.registerX)

		// case *ldaImmediate.opCode,
		// 	*ldaZeroPage.opCode,
		// 	*ldaZeroPageX.opCode,
		// 	*ldaAbsolute.opCode,
		// 	*ldaAbsoluteX.opCode,
		// 	*ldaAbsoluteY.opCode,
		// 	*ldaIndexedIndirect.opCode,
		// 	*ldaIndirectIndexed.opCode:
		// 	c.LDA(c.getAddrMode(op))

		// case LDX_immediate,
		// 	LDX_zeroPage,
		// 	LDX_zeroPageY,
		// 	LDX_absolute,
		// 	LDX_absoluteY:
		// 	c.LDX(c.getAddrMode(op))

		// case LDY_immediate,
		// 	LDY_zeroPage,
		// 	LDY_zeroPageX,
		// 	LDY_absolute,
		// 	LDY_absoluteX:
		// 	c.LDY(c.getAddrMode(op))

		// case STA_zeroPage,
		// 	STA_zeroPageX,
		// 	STA_absolute,
		// 	STA_absoluteX,
		// 	STA_absoluteY,
		// 	STA_indexedIndirect,
		// 	STA_indirectIndexed:
		// 	c.STA(c.getAddrMode(op))

		// case STX_zeroPage,
		// 	STX_zeroPageY,
		// 	STX_absolute:
		// 	c.STX(c.getAddrMode(op))

		// case INX:
		// 	c.INX()
		// 	c.updateZNFlags(c.registerX)
		// }
	}
}

func (c *cpu) loadAndRun(program []uint8) {
	c.load(program)
	c.reset()
	c.run()
}

func (c *cpu) TAX() {
	c.registerX = c.registerA
}

// func (c *cpu) LDA(mode addrMode) {
// 	addr := mode(c)
// 	param := c.memRead(addr)
// 	c.programCounter += 1
// 	c.registerA = param
// 	c.updateZNFlags(c.registerA)
// }

// func (c *cpu) LDX(mode addrMode) {
// 	addr := mode(c)
// 	param := c.memRead(addr)
// 	c.programCounter += 1
// 	c.registerX = param
// 	c.updateZNFlags(c.registerX)
// }

// func (c *cpu) LDY(mode addrMode) {
// 	addr := mode(c)
// 	param := c.memRead(addr)
// 	c.programCounter += 1
// 	c.registerY = param
// 	c.updateZNFlags(c.registerY)
// }

func (c *cpu) STA(mode addrMode) {
	addr := mode(c)
	c.memory[addr] = c.registerA
}

func (c *cpu) STX(mode addrMode) {
	addr := mode(c)
	c.memory[addr] = c.registerX
}

func (c *cpu) INX() {
	c.registerX += 1
}

func (c *cpu) updateZNFlags(result uint8) {
	if result == 0 {
		c.status |= uint8(1 << 1)
	} else {
		c.status &= ^uint8(1 << 1)
	}

	if result&uint8(1<<7) != 0 {
		c.status |= uint8(1 << 7)
	} else {
		c.status &= ^uint8(1 << 7)
	}
}

type addrMode func(*cpu) uint16

// func (c *cpu) getAddrMode(op uint8) addrMode {
// 	switch op {
// 	case LDA_immediate,
// 		LDX_immediate,
// 		LDY_immediate:
// 		return addrModeImmediate

// 	case LDA_zeroPage,
// 		LDX_zeroPage,
// 		LDY_zeroPage,
// 		STA_zeroPage:
// 		return addrModeZeroPage

// 	case LDA_zeroPageX,
// 		LDY_zeroPageX,
// 		STA_zeroPageX:
// 		return addrModeZeroPageX

// 	case LDX_zeroPageY:
// 		return addrModeZeroPageY

// 	case LDA_absolute,
// 		LDX_absolute,
// 		LDY_absolute,
// 		STA_absolute:
// 		return addrModeAbsolute

// 	case LDA_absoluteX,
// 		LDY_absoluteX,
// 		STA_absoluteX:
// 		return addrModeAbsoluteX

// 	case LDA_absoluteY,
// 		LDX_absoluteY,
// 		STA_absoluteY:
// 		return addrModeAbsoluteY

// 	case LDA_indexedIndirect,
// 		STA_indexedIndirect:
// 		return addrModeIndexedIndirect

// 	case LDA_indirectIndexed,
// 		STA_indirectIndexed:
// 		return addrModeIndirectIndexed

// 	default:
// 		log.Fatal("no addr mode found")
// 		return nil
// 	}
// }

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

func wrappingSumUint8(x, y uint8) uint8 {
	sum := x + y
	if sum > 0xFF {
		return sum - 0xFF
	}
	return sum
}

func wrappingSumUint16(x, y uint16) uint16 {
	sum := x + y
	if sum > 0xFFFF {
		return sum - 0xFFFF
	}
	return sum
}

func main() {}
