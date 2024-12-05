package operations

const (
	None = iota
	Multiply
	On
	Off
)

type Operation struct {
	Params        []int
	OperationType int
}
