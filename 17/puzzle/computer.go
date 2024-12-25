package puzzle

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	ADV int = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type Computer struct {
	A, B, C      int // register
	Instructions []int
	output       []int
}

func parseRegisters(regStr string) (int, error) {
	parts := strings.Split(regStr, ": ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid register format: %s", regStr)
	}
	return strconv.Atoi(strings.TrimSpace(parts[1]))
}

func parseInstructions(instStr string) ([]int, error) {
	parts := strings.Split(instStr, ": ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid instruction format: %s", instStr)
	}
	instParts := strings.Split(strings.TrimSpace(parts[1]), ",")
	instructions := make([]int, len(instParts))
	for i, val := range instParts {
		var err error
		instructions[i], err = strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			return nil, err
		}
	}
	return instructions, nil
}

func NewComputer(input []string) Computer {
	a, _ := parseRegisters(input[0])
	b, _ := parseRegisters(input[1])
	c, _ := parseRegisters(input[2])
	instructions, _ := parseInstructions(input[4])

	computer := Computer{
		A:            a,
		B:            b,
		C:            c,
		Instructions: instructions,
	}
	return computer
}

/*
The value of a combo operand can be found as follows:

Combo operands 0 through 3 represent literal values 0 through 3.
Combo operand 4 represents the value of register A.
Combo operand 5 represents the value of register B.
Combo operand 6 represents the value of register C.
Combo operand 7 is reserved and will not appear in valid programs.
*/
func (c *Computer) Combo(input int) int {
	switch input {
	case 0, 1, 2, 3:
		return input
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	case 7:
		fmt.Printf("Invalid combo operand\n")
	}

	return 0
}

func (c *Computer) division(input int) int {
	return int(float64(c.A) / math.Pow(2, float64(c.Combo(input))))
}

/*
The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an
operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of
the division operation is truncated to an integer and then written to the A register.
*/
func (c *Computer) Adv(input int) {
	c.A = c.division(input)
}

/*
The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's
literal operand, then stores the result in register B.
*/
func (c *Computer) Bxl(input int) {
	c.B = c.B ^ input
}

/*
The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby
keeping only its lowest 3 bits), then writes that value to the B register.
*/
func (c *Computer) Bst(input int) {
	c.B = c.Combo(input) % 8
}

/*
The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A
register is not zero, it jumps by setting the instruction pointer to the value of its
literal operand; if this instruction jumps, the instruction pointer is not increased
by 2 after this instruction.
*/
func (c *Computer) Jnz(input int) int {
	return input
}

/*
The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C,
then stores the result in register B. (For legacy reasons, this instruction reads an
operand but ignores it.)
*/
func (c *Computer) Bxc(_ int) {
	c.B = c.B ^ c.C
}

/*
The out instruction (opcode 5) calculates the value of its combo operand modulo 8,
then outputs that value. (If a program outputs multiple values, they are separated
by commas.)
*/
func (c *Computer) Out(input int) {
	c.output = append(c.output, c.Combo(input)%8)
}

/*
The bdv instruction (opcode 6) works exactly like the adv instruction except that
the result is stored in the B register. (The numerator is still read from the A register.)
*/
func (c *Computer) Bdv(input int) {
	c.B = c.division(input)
}

func (c *Computer) Cdv(input int) {
	c.C = c.division(input)
}

func (c *Computer) Execute(a, b int) {
	switch a {
	case ADV:
		c.Adv(b)
	case BXL:
		c.Bxl(b)
	case BST:
		c.Bst(b)
	case BXC:
		c.Bxc(b)
	case OUT:
		c.Out(b)
	case BDV:
		c.Bdv(b)
	case CDV:
		c.Cdv(b)
	}
}

func (c *Computer) Run() {
	for i := 0; i < len(c.Instructions); i += 2 {
		c.Execute(c.Instructions[i], c.Instructions[i+1])

		if c.Instructions[i] == JNZ && c.A != 0 {
			i = c.Instructions[i+1] - 2
		}
	}
}

func (c Computer) Output() string {
	strSlice := make([]string, len(c.output))
	for i, num := range c.output {
		strSlice[i] = strconv.Itoa(num)
	}

	return strings.Join(strSlice, ",")
}
