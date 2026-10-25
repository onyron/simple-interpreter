package cmpl

type constant struct {
	Value int `json:"CONST"`
}
type store struct {
	Identifier string `json:"ST"`
}

type load struct {
	Identifier string `json:"LD"`
}

type label struct {
	Name string `json:"LABEL"`
}

type jump struct {
	Label string `json:"JMP"`
}

type jnz struct {
	Label string `json:"JNZ"`
}

type binop struct {
	Op string `json:"BINOP"`
}
