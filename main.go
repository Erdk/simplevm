package main

import "fmt"

var R [16]int   // general purpose registers
var SP = &R[14] // stack pointer, R14
var PC = &R[15] // program counter, R15

/*
flag register
bit 0: Zero Flag
  1 - result of last comparison op is 0
  0 - reslut of last comparison op is different than 0
bit 1: Carry Flag
  1 - carry in last op
  0 - last op doesn't caused carry
*/
var FR int

/*
control registers

0x100 - 0x107 - non-maskable interrupts
0x108 - 0x10F - maskable interrupts
0x110 - intarrupts control register
bit 0:
  0 - maskable interrupts are ignored
  1 - maskable interrupts are not ignored
*/

/*
Ops

General layout:
opcode:                     1 byte
dst register (if required): 1 byte
src register (if required): 1 byte
imm(8/16/32) (if required): immediate value (1, 2, 4 bytes)

opcode (hex) | mnemonic
               description
				      usage

Memory access instructions:

00 | vmov rdst rsrc
     move from rsrc to rdst

		 vmov R2 R5 | 00 02 05

01 | vset rdst imm32
     set rdst to immediate value imm32

		 vset R4 0x1234 | 01 04 34 12 00 00

02 | vld rdst rsrc
     load 32 bits from address set in rsrc

		 vset R2 0x1234 | 01 02 34 12 00 00
     vld R1 R2      | 02 01 02
     this will load bytes from memory regions: 0x1234 0x1235 0x1236 0x1237

03 | vst rdst rsrc
     copy data from register rsrc to memory address stored in rdst

		 vset R9 0x1234     | 01 09 34 12 00 00
     vset R5 0x12345678 | 01 05 78 56 34 12
		 vst  R9 R5         | 03 09 05

		 Memory:
		 0x1234 78
		 0x1235 56
		 0x1236 34
		 0x1237 12

04 | vldb rdst rsrc
     load byte from memory address rsrc into register rdst

		 vset R3 0x1234 | 01 03 34 12 00 00
		 vldb R1 R3     | 04 01 03

05 | vstb rdst rsrc
     store lower byte of register rsrc to memory address from rdst

		 Set 0x41 to memory address 0x1234:

		 vset R1 0x41   | 01 01 41 00 00 00
		 vset R2 0x1234 | 01 02 34 12 00 00
		 vstb R2 R1     | 05 02 01

Arithmetic and logic instructions:

10 | vadd rdst rsrc
     sums immediate values from rdst and rsrc and stores result in rdst

		 Ex: add 5 from R1 and 8 from R2 and store result in R1
		 vset R1 0x5 | 01 01 05 00 00 00
		 vset R2 0x8 | 01 02 08 00 00 00
		 vadd R1 R2  | 10 01 02

11 | vsub rdst rsrc
     subtract rsrc form rdst, store result in rdst

		 Ex: subtract 5 in R1 from 13 in R2 and store result in R2
		 vset R1 0x5 | 01 01 05 00 00 00
		 vset R2 0xC | 01 02 13 00 00 00
		 vsub R2 R1  | 11 02 01

12 | vmul
13 | vdiv
14 | vmod
15 | vor
16 | vand
17 | vxor
18 | vnot
19 | vshl
1A | vshr

*/
func main() {
	fmt.Println("Hello world")
}
