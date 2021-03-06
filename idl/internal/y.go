// @generated Code generated by yacc

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

//line thrift.y:2
package internal

import __yyfmt__ "fmt"

//line thrift.y:2
import "go.uber.org/thriftrw/ast"

//line thrift.y:7
type yySymType struct {
	yys int
	// Used to record line numbers when the line number at the start point is
	// required.
	line int

	// Holds the final AST for the file.
	prog *ast.Program

	// Other intermediate variables:

	bul bool
	str string
	i64 int64
	dub float64

	fieldType     ast.Type
	structType    ast.StructureType
	baseTypeID    ast.BaseTypeID
	fieldRequired ast.Requiredness

	field  *ast.Field
	fields []*ast.Field

	header  ast.Header
	headers []ast.Header

	function  *ast.Function
	functions []*ast.Function

	enumItem  *ast.EnumItem
	enumItems []*ast.EnumItem

	definition  ast.Definition
	definitions []ast.Definition

	typeAnnotations []*ast.Annotation

	constantValue    ast.ConstantValue
	constantValues   []ast.ConstantValue
	constantMapItems []ast.ConstantMapItem
}

const IDENTIFIER = 57346
const LITERAL = 57347
const INTCONSTANT = 57348
const DUBCONSTANT = 57349
const NAMESPACE = 57350
const INCLUDE = 57351
const VOID = 57352
const BOOL = 57353
const BYTE = 57354
const I8 = 57355
const I16 = 57356
const I32 = 57357
const I64 = 57358
const DOUBLE = 57359
const STRING = 57360
const BINARY = 57361
const MAP = 57362
const LIST = 57363
const SET = 57364
const ONEWAY = 57365
const TYPEDEF = 57366
const STRUCT = 57367
const UNION = 57368
const EXCEPTION = 57369
const EXTENDS = 57370
const THROWS = 57371
const SERVICE = 57372
const ENUM = 57373
const CONST = 57374
const REQUIRED = 57375
const OPTIONAL = 57376
const TRUE = 57377
const FALSE = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"LITERAL",
	"INTCONSTANT",
	"DUBCONSTANT",
	"NAMESPACE",
	"INCLUDE",
	"VOID",
	"BOOL",
	"BYTE",
	"I8",
	"I16",
	"I32",
	"I64",
	"DOUBLE",
	"STRING",
	"BINARY",
	"MAP",
	"LIST",
	"SET",
	"ONEWAY",
	"TYPEDEF",
	"STRUCT",
	"UNION",
	"EXCEPTION",
	"EXTENDS",
	"THROWS",
	"SERVICE",
	"ENUM",
	"CONST",
	"REQUIRED",
	"OPTIONAL",
	"TRUE",
	"FALSE",
	"'*'",
	"'='",
	"'{'",
	"'}'",
	"':'",
	"'('",
	"')'",
	"'<'",
	"','",
	"'>'",
	"'['",
	"']'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	8, 70,
	9, 70,
	-2, 8,
	-1, 3,
	1, 1,
	-2, 70,
}

const yyNprod = 74
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 179

var yyAct = [...]int{

	26, 87, 55, 5, 7, 10, 63, 62, 65, 25,
	70, 66, 67, 11, 122, 124, 94, 12, 93, 92,
	155, 11, 59, 58, 27, 12, 57, 146, 145, 114,
	90, 153, 56, 56, 56, 137, 89, 139, 120, 85,
	68, 69, 115, 85, 88, 70, 66, 67, 79, 76,
	53, 106, 60, 118, 64, 71, 130, 51, 50, 54,
	82, 52, 78, 81, 150, 24, 104, 73, 74, 75,
	127, 128, 125, 91, 100, 68, 69, 9, 8, 95,
	22, 21, 98, 96, 133, 101, 99, 31, 141, 102,
	129, 109, 97, 86, 105, 49, 111, 112, 23, 34,
	113, 33, 110, 32, 116, 30, 29, 28, 71, 121,
	149, 103, 84, 117, 72, 123, 119, 108, 107, 3,
	6, 61, 71, 77, 83, 131, 2, 134, 135, 4,
	132, 80, 16, 138, 126, 35, 136, 1, 0, 140,
	71, 0, 0, 81, 144, 142, 71, 0, 143, 147,
	0, 151, 152, 0, 148, 81, 14, 18, 19, 20,
	39, 154, 17, 15, 13, 0, 0, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 36, 37, 38,
}
var yyPact = [...]int{

	-1000, -1000, -1000, -1000, -1000, 69, -32, 132, 76, 61,
	-1000, -1000, -1000, -1000, -1000, 103, 102, 101, -1000, -1000,
	-1000, -1000, 82, 99, 97, 95, 156, 91, 19, 18,
	22, -1000, -1000, -1000, 21, -8, -18, -21, -22, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -8,
	-1000, -1000, -1000, -1000, 40, -1000, -1000, -1000, -1000, -1000,
	-1000, 9, 8, 20, 89, -1000, -1000, -1000, -1000, -1000,
	-1000, -3, -13, -26, -28, -30, -8, -32, 88, -8,
	-32, 68, -8, -32, 56, -1000, 12, -1000, -1000, -1000,
	-1000, 87, -1000, -8, -8, -1000, -1000, -9, -1000, -1000,
	1, -1000, -1000, -1000, -1000, -1000, -1000, 5, -2, -24,
	-31, -1000, -1000, -1000, 66, 37, 86, 16, -1000, -32,
	-1000, 40, 79, -1000, -8, -8, -1000, -1000, -1000, -7,
	-8, -1000, -4, -32, -1000, -1000, 84, -1000, -1000, 40,
	-1000, -10, -16, -32, -1000, 40, 35, -1000, -8, -8,
	-11, -1000, -1000, -1000, -23, -1000,
}
var yyPgo = [...]int{

	0, 0, 137, 9, 135, 134, 132, 131, 7, 129,
	126, 124, 6, 123, 121, 120, 119, 8, 118, 117,
	114, 2, 5, 112, 111, 110,
}
var yyR1 = [...]int{

	0, 2, 10, 10, 9, 9, 9, 9, 16, 16,
	15, 15, 15, 15, 15, 15, 6, 6, 6, 14,
	14, 13, 13, 8, 8, 7, 7, 5, 5, 5,
	12, 12, 11, 23, 23, 24, 24, 25, 25, 3,
	3, 3, 3, 3, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 17, 17, 17, 17, 17, 17, 17,
	17, 18, 18, 19, 19, 21, 21, 20, 20, 20,
	1, 22, 22, 22,
}
var yyR2 = [...]int{

	0, 2, 0, 2, 3, 4, 4, 4, 0, 3,
	6, 5, 7, 7, 7, 10, 1, 1, 1, 0,
	3, 3, 5, 0, 3, 7, 9, 1, 1, 0,
	0, 3, 9, 1, 0, 1, 1, 0, 4, 3,
	8, 6, 6, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	4, 0, 3, 0, 6, 0, 3, 0, 6, 4,
	0, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -2, -10, -16, -9, -1, -15, -1, 9, 8,
	-22, 45, 49, 32, 24, 31, -6, 30, 25, 26,
	27, 5, 4, 37, 4, -3, -1, -3, 4, 4,
	4, 5, 4, 4, 4, -4, 20, 21, 22, 4,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 4,
	39, 39, 39, 28, 38, -21, 42, 44, 44, 44,
	-21, -14, -8, -12, -1, -17, 6, 7, 35, 36,
	5, -1, -20, -3, -3, -3, 40, -13, -1, 40,
	-7, -1, 40, -11, -23, 23, 4, 4, 47, 39,
	43, -1, 45, 46, 46, -21, -22, 4, -21, -22,
	6, -21, -22, -24, 10, -3, 39, -18, -19, 4,
	-3, -21, -21, -21, 38, 41, -1, -12, 48, -17,
	40, -1, 38, -22, 46, 6, -5, 33, 34, 4,
	40, -22, -17, 5, -21, -21, -3, 42, -21, 41,
	-22, 4, -8, -17, -21, 38, 43, -22, -17, -25,
	29, -21, -21, 42, -8, 43,
}
var yyDef = [...]int{

	2, -2, -2, -2, 3, 0, 73, 0, 0, 0,
	9, 71, 72, 70, 70, 0, 0, 0, 16, 17,
	18, 4, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 5, 6, 7, 0, 65, 0, 0, 0, 43,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 65,
	19, 23, 30, 70, 70, 39, 67, 70, 70, 70,
	11, 70, 70, 34, 0, 10, 53, 54, 55, 56,
	57, 0, 70, 0, 0, 0, 65, 73, 0, 65,
	73, 0, 65, 73, 70, 33, 0, 58, 61, 63,
	66, 0, 70, 65, 65, 12, 20, 65, 13, 24,
	0, 14, 31, 70, 35, 36, 30, 70, 70, 73,
	0, 41, 42, 21, 0, 29, 0, 34, 59, 73,
	60, 70, 0, 69, 65, 65, 70, 27, 28, 0,
	65, 62, 0, 73, 40, 22, 0, 23, 15, 70,
	68, 65, 70, 73, 25, 70, 37, 64, 65, 65,
	0, 26, 32, 23, 70, 38,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	42, 43, 37, 3, 45, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 49,
	44, 38, 46, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 47, 3, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 39, 3, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:92
		{
			yyVAL.prog = &ast.Program{Headers: yyDollar[1].headers, Definitions: yyDollar[2].definitions}
			yylex.(*lexer).program = yyVAL.prog
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:104
		{
			yyVAL.headers = nil
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:105
		{
			yyVAL.headers = append(yyDollar[1].headers, yyDollar[2].header)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:110
		{
			yyVAL.header = &ast.Include{
				Path: yyDollar[3].str,
				Line: yyDollar[1].line,
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:117
		{
			yyVAL.header = &ast.Include{
				Name: yyDollar[3].str,
				Path: yyDollar[4].str,
				Line: yyDollar[1].line,
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:125
		{
			yyVAL.header = &ast.Namespace{
				Scope: "*",
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:133
		{
			yyVAL.header = &ast.Namespace{
				Scope: yyDollar[3].str,
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:147
		{
			yyVAL.definitions = nil
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:148
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:155
		{
			yyVAL.definition = &ast.Constant{
				Name:  yyDollar[4].str,
				Type:  yyDollar[3].fieldType,
				Value: yyDollar[6].constantValue,
				Line:  yyDollar[1].line,
			}
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:165
		{
			yyVAL.definition = &ast.Typedef{
				Name:        yyDollar[4].str,
				Type:        yyDollar[3].fieldType,
				Annotations: yyDollar[5].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:174
		{
			yyVAL.definition = &ast.Enum{
				Name:        yyDollar[3].str,
				Items:       yyDollar[5].enumItems,
				Annotations: yyDollar[7].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 13:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:183
		{
			yyVAL.definition = &ast.Struct{
				Name:        yyDollar[3].str,
				Type:        yyDollar[2].structType,
				Fields:      yyDollar[5].fields,
				Annotations: yyDollar[7].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 14:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:194
		{
			yyVAL.definition = &ast.Service{
				Name:        yyDollar[3].str,
				Functions:   yyDollar[5].functions,
				Annotations: yyDollar[7].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 15:
		yyDollar = yyS[yypt-10 : yypt+1]
		//line thrift.y:204
		{
			parent := &ast.ServiceReference{
				Name: yyDollar[6].str,
				Line: yyDollar[5].line,
			}

			yyVAL.definition = &ast.Service{
				Name:        yyDollar[3].str,
				Functions:   yyDollar[8].functions,
				Parent:      parent,
				Annotations: yyDollar[10].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:221
		{
			yyVAL.structType = ast.StructType
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:222
		{
			yyVAL.structType = ast.UnionType
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:223
		{
			yyVAL.structType = ast.ExceptionType
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:227
		{
			yyVAL.enumItems = nil
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:228
		{
			yyVAL.enumItems = append(yyDollar[1].enumItems, yyDollar[2].enumItem)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:233
		{
			yyVAL.enumItem = &ast.EnumItem{Name: yyDollar[2].str, Annotations: yyDollar[3].typeAnnotations, Line: yyDollar[1].line}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line thrift.y:235
		{
			value := int(yyDollar[4].i64)
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[2].str,
				Value:       &value,
				Annotations: yyDollar[5].typeAnnotations,
				Line:        yyDollar[1].line,
			}
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:247
		{
			yyVAL.fields = nil
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:248
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[2].field)
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line thrift.y:254
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[2].i64),
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Annotations:  yyDollar[7].typeAnnotations,
				Line:         yyDollar[1].line,
			}
		}
	case 26:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line thrift.y:266
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[2].i64),
				Name:         yyDollar[6].str,
				Type:         yyDollar[5].fieldType,
				Requiredness: yyDollar[4].fieldRequired,
				Default:      yyDollar[8].constantValue,
				Annotations:  yyDollar[9].typeAnnotations,
				Line:         yyDollar[1].line,
			}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:280
		{
			yyVAL.fieldRequired = ast.Required
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:281
		{
			yyVAL.fieldRequired = ast.Optional
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:282
		{
			yyVAL.fieldRequired = ast.Unspecified
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:286
		{
			yyVAL.functions = nil
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:287
		{
			yyVAL.functions = append(yyDollar[1].functions, yyDollar[2].function)
		}
	case 32:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line thrift.y:293
		{
			yyVAL.function = &ast.Function{
				Name:        yyDollar[4].str,
				Parameters:  yyDollar[6].fields,
				ReturnType:  yyDollar[2].fieldType,
				Exceptions:  yyDollar[8].fields,
				OneWay:      yyDollar[1].bul,
				Annotations: yyDollar[9].typeAnnotations,
				Line:        yyDollar[3].line,
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:307
		{
			yyVAL.bul = true
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:308
		{
			yyVAL.bul = false
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:312
		{
			yyVAL.fieldType = nil
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:313
		{
			yyVAL.fieldType = yyDollar[1].fieldType
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:317
		{
			yyVAL.fields = nil
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:318
		{
			yyVAL.fields = yyDollar[3].fields
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:327
		{
			yyVAL.fieldType = ast.BaseType{ID: yyDollar[2].baseTypeID, Annotations: yyDollar[3].typeAnnotations, Line: yyDollar[1].line}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line thrift.y:331
		{
			yyVAL.fieldType = ast.MapType{KeyType: yyDollar[4].fieldType, ValueType: yyDollar[6].fieldType, Annotations: yyDollar[8].typeAnnotations, Line: yyDollar[1].line}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:333
		{
			yyVAL.fieldType = ast.ListType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:335
		{
			yyVAL.fieldType = ast.SetType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:337
		{
			yyVAL.fieldType = ast.TypeReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:341
		{
			yyVAL.baseTypeID = ast.BoolTypeID
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:342
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:343
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:344
		{
			yyVAL.baseTypeID = ast.I16TypeID
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:345
		{
			yyVAL.baseTypeID = ast.I32TypeID
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:346
		{
			yyVAL.baseTypeID = ast.I64TypeID
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:347
		{
			yyVAL.baseTypeID = ast.DoubleTypeID
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:348
		{
			yyVAL.baseTypeID = ast.StringTypeID
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:349
		{
			yyVAL.baseTypeID = ast.BinaryTypeID
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:357
		{
			yyVAL.constantValue = ast.ConstantInteger(yyDollar[1].i64)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:358
		{
			yyVAL.constantValue = ast.ConstantDouble(yyDollar[1].dub)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:359
		{
			yyVAL.constantValue = ast.ConstantBoolean(true)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:360
		{
			yyVAL.constantValue = ast.ConstantBoolean(false)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line thrift.y:361
		{
			yyVAL.constantValue = ast.ConstantString(yyDollar[1].str)
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line thrift.y:363
		{
			yyVAL.constantValue = ast.ConstantReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:365
		{
			yyVAL.constantValue = ast.ConstantList{Items: yyDollar[3].constantValues, Line: yyDollar[1].line}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:366
		{
			yyVAL.constantValue = ast.ConstantMap{Items: yyDollar[3].constantMapItems, Line: yyDollar[1].line}
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:370
		{
			yyVAL.constantValues = nil
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:372
		{
			yyVAL.constantValues = append(yyDollar[1].constantValues, yyDollar[2].constantValue)
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:376
		{
			yyVAL.constantMapItems = nil
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:378
		{
			yyVAL.constantMapItems = append(yyDollar[1].constantMapItems, ast.ConstantMapItem{Key: yyDollar[3].constantValue, Value: yyDollar[5].constantValue, Line: yyDollar[2].line})
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:386
		{
			yyVAL.typeAnnotations = nil
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line thrift.y:387
		{
			yyVAL.typeAnnotations = yyDollar[2].typeAnnotations
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:391
		{
			yyVAL.typeAnnotations = nil
		}
	case 68:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line thrift.y:393
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Value: yyDollar[5].str, Line: yyDollar[2].line})
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line thrift.y:395
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Line: yyDollar[2].line})
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line thrift.y:412
		{
			yyVAL.line = yylex.(*lexer).line
		}
	}
	goto yystack /* stack new state and value */
}
