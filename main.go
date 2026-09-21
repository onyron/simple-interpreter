package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var vars = map[string]int{}

func doBinOp(op string, left int, right int) int {
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

func parseExpr(expr map[string]any) int {
	for k := range expr {
		if k == "var" {
			return vars[expr[k].(string)]
		}
		if k == "const" {
			return int(expr[k].(float64))
		}
		if k == "binop" {
			left := parseExpr(expr["left"].(map[string]any))
			right := parseExpr(expr["right"].(map[string]any))
			return doBinOp(expr["binop"].(string), left, right)
		}
	}
	panic("unreachable")
}

func handleStmt(anyStmt any) {
	stmt, ok := anyStmt.(string)
	if ok {
		parseStmt(stmt)
	} else {
		parseStmtMap(anyStmt.(map[string]any))
	}
}

func parseSeq(seq map[string]any) {
	handleStmt(seq["left"])
	handleStmt(seq["right"])
}

func parseRead(ident string) {
	var val int
	_, err := fmt.Scan(&val)
	if err != nil {
		panic(err)
	}
	vars[ident] = val
}

func parseWrite(expr map[string]any) {
	fmt.Println(parseExpr(expr))
}

func parseAssn(assn map[string]any) {
	src := parseExpr(assn["src"].(map[string]any))
	vars[assn["dst"].(string)] = src
}

func parseWhile(while map[string]any) {
	cond := parseExpr(while["cond"].(map[string]any))
	for cond != 0 {
		handleStmt(while["body"])
		cond = parseExpr(while["cond"].(map[string]any))
	}
}

func parseDo(do map[string]any) {
	handleStmt(do["body"])
	cond := parseExpr(do["cond"].(map[string]any))
	for cond != 0 {
		handleStmt(do["body"])
		cond = parseExpr(do["cond"].(map[string]any))
	}
}

func parseIf(ifStmt map[string]any) {
	cond := parseExpr(ifStmt["cond"].(map[string]any))
	if cond != 0 {
		handleStmt(ifStmt["then"])
	} else {
		handleStmt(ifStmt["else"])
	}
}

func parseStmt(stmt string) {
	if stmt == "skip" {
		return
	}

	panic("unknown statement: " + stmt)
}

func parseStmtMap(stmt map[string]any) {
	for k := range stmt {
		if k == "seq" {
			parseSeq(stmt["seq"].(map[string]any))
		} else if k == "read" {
			parseRead(stmt["read"].(string))
		} else if k == "write" {
			parseWrite(stmt["write"].(map[string]any))
		} else if k == "assn" {
			parseAssn(stmt["assn"].(map[string]any))
		} else if k == "while" {
			parseWhile(stmt["while"].(map[string]any))
		} else if k == "do" {
			parseDo(stmt["do"].(map[string]any))
		} else if k == "if" {
			parseIf(stmt["if"].(map[string]any))
		} else {
			panic("unknown statement: " + k)
		}
	}
}

func readJSON(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}

func parseJSON(data []byte) any {
	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	return result
}

func main() {
	data := readJSON("contents.json")
	result := parseJSON(data)
	handleStmt(result)
}
