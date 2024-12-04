package operations

func mul(a, b int) int {
	return a * b
}

func Execute(ops Operation) int {
	switch ops.OperationType {
	case Multiply:
		return mul(ops.Params[0], ops.Params[1])
	}
	return 0
}
