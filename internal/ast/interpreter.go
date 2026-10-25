package ast

import (
	"encoding/json"
	"fmt"
	"os"
)

type Interpreter struct {
	vars map[string]int
}

func New() *Interpreter {
	return &Interpreter{
		vars: make(map[string]int),
	}
}

func (in *Interpreter) Run(program any) {
	in.handleStmt(program)
}

func (in *Interpreter) doBinOp(op string, left int, right int) int {
	if op == "!!" {
		if left == 1 || right == 1 {
			return 1
		}
		return 0
	} else if op == "&&" {
		if left == 0 || right == 0 {
			return 0
		}
		return 1
	} else if op == "==" {
		if left == right {
			return 1
		}
		return 0
	} else if op == "!=" {
		if left != right {
			return 1
		}
		return 0
	} else if op == "<=" {
		if left <= right {
			return 1
		}
		return 0
	} else if op == "<" {
		if left < right {
			return 1
		}
		return 0
	} else if op == ">=" {
		if left >= right {
			return 1
		}
		return 0
	} else if op == ">" {
		if left > right {
			return 1
		}
		return 0
	} else if op == "+" {
		return left + right
	} else if op == "-" {
		return left - right
	} else if op == "*" {
		return left * right
	} else if op == "/" {
		return left / right
	} else if op == "%" {
		return left % right
	}

	panic("invalid operation: " + op)
}

func (in *Interpreter) parseExpr(expr map[string]any) int {
	for k := range expr {
		if k == "var" {
			return in.vars[expr[k].(string)]
		}
		if k == "const" {
			return int(expr[k].(float64))
		}
		if k == "binop" {
			left := in.parseExpr(expr["left"].(map[string]any))
			right := in.parseExpr(expr["right"].(map[string]any))
			return in.doBinOp(expr["binop"].(string), left, right)
		}
	}
	panic("unreachable")
}

func (in *Interpreter) handleStmt(anyStmt any) {
	stmt, ok := anyStmt.(string)
	if ok {
		in.parseStmt(stmt)
	} else {
		in.parseStmtMap(anyStmt.(map[string]any))
	}
}

func (in *Interpreter) parseSeq(seq map[string]any) {
	in.handleStmt(seq["left"])
	in.handleStmt(seq["right"])
}

func (in *Interpreter) parseRead(ident string) {
	var val int
	_, err := fmt.Scan(&val)
	if err != nil {
		panic(err)
	}
	in.vars[ident] = val
}

func (in *Interpreter) parseWrite(expr map[string]any) {
	fmt.Println(in.parseExpr(expr))
}

func (in *Interpreter) parseAssn(assn map[string]any) {
	src := in.parseExpr(assn["src"].(map[string]any))
	in.vars[assn["dst"].(string)] = src
}

func (in *Interpreter) parseWhile(while map[string]any) {
	cond := in.parseExpr(while["cond"].(map[string]any))
	for cond != 0 {
		in.handleStmt(while["body"])
		cond = in.parseExpr(while["cond"].(map[string]any))
	}
}

func (in *Interpreter) parseDo(do map[string]any) {
	in.handleStmt(do["body"])
	cond := in.parseExpr(do["cond"].(map[string]any))
	for cond != 0 {
		in.handleStmt(do["body"])
		cond = in.parseExpr(do["cond"].(map[string]any))
	}
}

func (in *Interpreter) parseIf(ifStmt map[string]any) {
	cond := in.parseExpr(ifStmt["cond"].(map[string]any))
	if cond != 0 {
		in.handleStmt(ifStmt["then"])
	} else {
		in.handleStmt(ifStmt["else"])
	}
}

func (in *Interpreter) parseStmt(stmt string) {
	if stmt == "skip" {
		return
	}

	panic("unknown statement: " + stmt)
}

func (in *Interpreter) parseStmtMap(stmt map[string]any) {
	for k := range stmt {
		if k == "seq" {
			in.parseSeq(stmt["seq"].(map[string]any))
		} else if k == "read" {
			in.parseRead(stmt["read"].(string))
		} else if k == "write" {
			in.parseWrite(stmt["write"].(map[string]any))
		} else if k == "assn" {
			in.parseAssn(stmt["assn"].(map[string]any))
		} else if k == "while" {
			in.parseWhile(stmt["while"].(map[string]any))
		} else if k == "do" {
			in.parseDo(stmt["do"].(map[string]any))
		} else if k == "if" {
			in.parseIf(stmt["if"].(map[string]any))
		} else {
			panic("unknown statement: " + k)
		}
	}
}

func LoadProgram(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
