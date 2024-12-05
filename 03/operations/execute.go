package operations

func mul(a, b int) int {
	return a * b
}

func ExecuteOperations(ops []Operation) int {
	enabled := true
	total := 0
	for _, op := range ops {
		switch op.OperationType {
		case Multiply:
			if enabled {
				total += mul(op.Params[0], op.Params[1])
			}

		case Enable:
			enabled = true

		case Disable:
			enabled = false
		}
	}
	return total
}
