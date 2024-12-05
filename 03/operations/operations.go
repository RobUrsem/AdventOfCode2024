package operations

const (
	None = iota
	Multiply
	Enable
	Disable
)

type Operation struct {
	Params        []int
	OperationType int
}
