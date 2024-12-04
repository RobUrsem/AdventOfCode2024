package operations

const (
	None = iota
	Multiply
)

type Operation struct {
	Params        []int
	OperationType int
}
