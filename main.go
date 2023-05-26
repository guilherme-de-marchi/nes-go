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
	c.setSTA()
	c.setSTX()
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
