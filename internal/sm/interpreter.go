package sm

type Interpreter struct {
	ip           int
	vars         map[string]int
	labels       map[string]int
	stack        Stack[int]
	instructions []any
}

func New(instructions []any) *Interpreter {
	vars := make(map[string]int)
	labels := make(map[string]int)
	stack := Stack[int]{}
	in := Interpreter{
		ip:           0,
		vars:         vars,
		labels:       labels,
		stack:        stack,
		instructions: instructions,
	}

	in.findLabels()
	return &in
}

func (in *Interpreter) Run() {
	in.handleInstructions()
}
