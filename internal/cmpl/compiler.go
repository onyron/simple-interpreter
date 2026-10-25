package cmpl

import (
	"strconv"
)

type Compiler struct {
	instructions []any
	labelCnt     int
}

func New() *Compiler {
	return &Compiler{}
}

func (c *Compiler) Compile(program any) []any {
	c.handleStmt(program)
	return c.instructions
}

func (c *Compiler) parseExpr(expr map[string]any) {
	for k := range expr {
		if k == "var" {
			ld := load{Identifier: expr[k].(string)}
			c.instructions = append(c.instructions, ld)
			return
		}
		if k == "const" {
			value := int(expr[k].(float64))
			constExpr := constant{value}
			c.instructions = append(c.instructions, constExpr)
			return
		}
		if k == "binop" {
			c.parseExpr(expr["left"].(map[string]any))
			c.parseExpr(expr["right"].(map[string]any))
			op := binop{expr["binop"].(string)}
			c.instructions = append(c.instructions, op)
			return
		}
	}
	panic("unreachable")
}

func (c *Compiler) handleStmt(anyStmt any) {
	stmt, ok := anyStmt.(string)
	if ok {
		c.parseStmt(stmt)
	} else {
		c.parseStmtMap(anyStmt.(map[string]any))
	}
}

func (c *Compiler) parseSeq(seq map[string]any) {
	c.handleStmt(seq["left"])
	c.handleStmt(seq["right"])
}

func (c *Compiler) parseRead(ident string) {
	st := store{Identifier: ident}
	c.instructions = append(c.instructions, "READ", st)
}

func (c *Compiler) parseWrite(expr map[string]any) {
	c.parseExpr(expr)
	c.instructions = append(c.instructions, "WRITE")
}

func (c *Compiler) parseAssn(assn map[string]any) {
	c.parseExpr(assn["src"].(map[string]any))
	st := store{Identifier: assn["dst"].(string)}
	c.instructions = append(c.instructions, st)
}

func (c *Compiler) parseWhile(while map[string]any) {
	cody := "L_while_cody_" + strconv.Itoa(c.labelCnt)
	c.labelCnt++
	cond := "L_while_cond_" + strconv.Itoa(c.labelCnt)
	c.labelCnt++
	codyLabel := label{cody}
	condLabel := label{cond}
	jmp := jump{cond}
	jumpnz := jnz{cody}
	c.instructions = append(c.instructions, jmp, codyLabel)
	c.handleStmt(while["body"])
	c.instructions = append(c.instructions, condLabel)
	c.parseExpr(while["cond"].(map[string]any))
	c.instructions = append(c.instructions, jumpnz)
}

func (c *Compiler) parseDo(do map[string]any) {
	cody := "L_while_cody_" + strconv.Itoa(c.labelCnt)
	c.labelCnt++
	codyLabel := label{cody}
	jumpnz := jnz{cody}
	c.instructions = append(c.instructions, codyLabel)
	c.handleStmt(do["body"])
	c.parseExpr(do["cond"].(map[string]any))
	c.instructions = append(c.instructions, jumpnz)
}

func (c *Compiler) parseIf(ifStmt map[string]any) {
	c.parseExpr(ifStmt["cond"].(map[string]any))
	end := "L_end_" + strconv.Itoa(c.labelCnt)
	c.labelCnt++
	elseStr := "L_else_" + strconv.Itoa(c.labelCnt)
	c.labelCnt++
	endLabel := label{end}
	elseLabel := label{elseStr}
	jmp := jump{end}
	jumpnz := jnz{elseStr}
	c.instructions = append(c.instructions, jumpnz)
	c.handleStmt(ifStmt["then"])
	c.instructions = append(c.instructions, jmp, elseLabel)
	c.handleStmt(ifStmt["else"])
	c.instructions = append(c.instructions, endLabel)
}

func (c *Compiler) parseStmt(stmt string) {
	if stmt == "skip" {
		return
	}

	panic("unknown statement: " + stmt)
}

func (c *Compiler) parseStmtMap(stmt map[string]any) {
	for k := range stmt {
		if k == "seq" {
			c.parseSeq(stmt["seq"].(map[string]any))
		} else if k == "write" {
			c.parseWrite(stmt["write"].(map[string]any))
		} else if k == "assn" {
			c.parseAssn(stmt["assn"].(map[string]any))
		} else if k == "while" {
			c.parseWhile(stmt["while"].(map[string]any))
		} else if k == "read" {
			c.parseRead(stmt["read"].(string))
		} else if k == "do" {
			c.parseDo(stmt["do"].(map[string]any))
		} else if k == "if" {
			c.parseIf(stmt["if"].(map[string]any))
		} else {
			panic("unknown statement: " + k)
		}
	}
}
