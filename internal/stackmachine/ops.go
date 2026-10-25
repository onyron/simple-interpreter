package stackmachine

func (in *Interpreter) doBinOp(op string) {
	r := in.stack.Pop()
	l := in.stack.Pop()
	in.stack.Push(parseBinOp(op, l, r))
}

func parseBinOp(op string, left, right int) int {
	switch op {
	case "!!":
		if left == 1 || right == 1 {
			return 1
		}
		return 0
	case "&&":
		if left == 0 || right == 0 {
			return 0
		}
		return 1
	case "==":
		if left == right {
			return 1
		}
		return 0
	case "!=":
		if left != right {
			return 1
		}
		return 0
	case "<=":
		if left <= right {
			return 1
		}
		return 0
	case "<":
		if left < right {
			return 1
		}
		return 0
	case ">=":
		if left >= right {
			return 1
		}
		return 0
	case ">":
		if left > right {
			return 1
		}
		return 0
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	case "%":
		return left % right
	default:
		panic("Unknown op: " + op)
	}
}
