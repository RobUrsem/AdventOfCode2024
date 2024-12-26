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
	Output       []int
	CorrectedA   int
	Expected     string
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
		CorrectedA:   math.MaxInt,
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

func (c *Computer) shr(input int) int {
	return c.A >> c.Combo(input)
}

/*
The adv instruction (opcode 0) performs division. The numerator is the value in the A register.
The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an
operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of
the division operation is truncated to an integer and then written to the A register.
*/
func (c *Computer) Adv(input int) {
	c.A = c.shr(input)
}

/*
The bdv instruction (opcode 6) works exactly like the adv instruction except that
the result is stored in the B register. (The numerator is still read from the A register.)
*/
func (c *Computer) Bdv(input int) {
	c.B = c.shr(input)
}

/*
The cdv instruction (opcode 7) works exactly like the adv instruction except that
the result is stored in the C register. (The numerator is still read from the A register.)
*/
func (c *Computer) Cdv(input int) {
	c.C = c.shr(input)
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
	c.Output = append(c.Output, c.Combo(input)%8)
}

func (c *Computer) Execute(op, arg int) {
	switch op {
	case ADV:
		c.Adv(arg)
	case BXL:
		c.Bxl(arg)
	case BST:
		c.Bst(arg)
	case BXC:
		c.Bxc(arg)
	case OUT:
		c.Out(arg)
	case BDV:
		c.Bdv(arg)
	case CDV:
		c.Cdv(arg)
	}
}

func (c *Computer) Run() string {
	c.Output = []int{}

	for i := 0; i < len(c.Instructions); i += 2 {
		op := c.Instructions[i]
		arg := c.Instructions[i+1]
		c.Execute(op, arg)

		if op == JNZ && c.A != 0 {
			i = arg - 2
		}
	}

	return c.Print()
}

func (c *Computer) Reset(regA, regB, regC int) {
	c.A = regA
	c.B = regB
	c.C = regC
	c.Output = []int{}
}

/*
-----------------------------------------------
Our program:
1: 2,4, B = A & 7
2: 1,1, B = B ^ 1
3: 7,5, C = A >> B
4: 0,3, A = A >> 3
5: 4,7, B = B ^ C
6: 1,6, B = B ^ 6
7: 5,5, print B
8: 3,0, if A == 0 then halt else goto start

Note that in instruction 4 we shift A by 3 bits
thus dividing by 8. We only need to look at 3 bits
at a time to get the next numbers.

To get the pattern that gets us the output that is
equal to that of the input instructions, we should
start at the end, and play with the last three bits
of the number. Once we found a suitable candidate
we take that, multiply it by 8 bits (<< 3) and try
for the next output number until we have found all.

Since there might be multiple solutions (modulo 8)
we need to take the lowest number.
------------------------------------------------
*/
func (c *Computer) RunReverse() {
	// Start at the end
	index := len(c.Instructions) - 1

	// Convert the instructions to a single string
	strSlice := make([]string, len(c.Instructions))
	for i, num := range c.Instructions {
		strSlice[i] = strconv.Itoa(num)
	}
	c.Expected = strings.Join(strSlice, "")

	// Let's go
	c.solve(0, index)
}

func (c *Computer) solve(currentA int, index int) {
	if index == -1 {
		c.CorrectedA = min(c.CorrectedA, currentA)
		return
	}

	// get the last (length - index) characters
	expected := c.Expected[index:]
	for remainder := 0; remainder < 8; remainder++ {
		nextA := currentA*8 + remainder

		c.Reset(nextA, 0, 0)
		output := c.Run()

		if output == expected {
			// get to the next number
			c.solve(nextA, index-1)
		}
	}
}

func (c Computer) Print() string {
	strSlice := make([]string, len(c.Output))
	for i, num := range c.Output {
		strSlice[i] = strconv.Itoa(num)
	}

	return strings.Join(strSlice, "")
}
