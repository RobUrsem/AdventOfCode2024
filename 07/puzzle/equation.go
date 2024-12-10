package puzzle

const (
	NOP = iota
	ADD
	MUL
	CAT
)

type Equation struct {
	Answer       int64
	Coefficients []int64
	Operators    []int
	Valid        bool
}

type Equations []Equation

func AreEqual(a, b Equation) bool {
	if a.Answer != b.Answer {
		return false
	}

	if len(a.Coefficients) != len(b.Coefficients) {
		return false
	}

	for i := 0; i < len(a.Coefficients); i++ {
		if a.Coefficients[i] != b.Coefficients[i] {
			return false
		}
	}

	return true
}
