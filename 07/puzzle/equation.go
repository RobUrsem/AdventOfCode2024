package puzzle

const (
	NOP = iota
	ADD
	MUL
)

type Equation struct {
	answer       int64
	coefficients []int64
	operators    []int
}

type Equations []Equation

func AreEqual(a, b Equation) bool {
	if a.answer != b.answer {
		return false
	}

	if len(a.coefficients) != len(b.coefficients) {
		return false
	}

	for i := 0; i < len(a.coefficients); i++ {
		if a.coefficients[i] != b.coefficients[i] {
			return false
		}
	}

	return true
}
