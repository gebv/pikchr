//line lexer.rl:1
package syntax

import (
	"fmt"
)

//line lexer.go:11
const md_parser_start int = 0
const md_parser_first_final int = 0
const md_parser_error int = -1

const md_parser_en_main int = 0

//line lexer.rl:21

type lexer struct {
	// It must be an array containting the data to process.
	data []byte

	// Data end pointer. This should be initialized to p plus the data length on every run of the machine. In Java and Ruby code this should be initialized to the data length.
	pe int

	// Data pointer. In C/D code this variable is expected to be a pointer to the character data to process. It should be initialized to the beginning of the data block on every run of the machine. In Java and Ruby it is used as an offset to data and must be an integer. In this case it should be initialized to zero on every run of the machine.
	p int

	// This must be a pointer to character data. In Java and Ruby code this must be an integer. See Section 6.3 for more information.
	ts int

	// Also a pointer to character data.
	te int

	// This must be an integer value. It is a variable sometimes used by scanner code to keep track of the most recent successful pattern match.
	act int

	// Current state. This must be an integer and it should persist across invocations of the machine when the data is broken into blocks that are processed independently. This variable may be modified from outside the execution loop, but not from within.
	cs int

	// This must be an integer value and will be used as an offset to stack, giving the next available spot on the top of the stack.
	top int

	result Blocks
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data:   data,
		pe:     len(data),
		result: Blocks{},
	}

//line lexer.go:59
	{
		lex.cs = md_parser_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line lexer.rl:60
	return lex
}

func (lex *lexer) Lex(out *mdSymType) int {
	eof := lex.pe
	tok := 0

//line lexer.go:75
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		switch lex.cs {
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		}
		goto st_out
	tr2:
//line lexer.rl:92
		lex.te = (lex.p) + 1
		{
			if mdDebug == 1 {
				fmt.Print("EOL")
			}

			out.token = eolToken{rawToken(string(lex.data[lex.ts:lex.te]))}
			tok = NEWLINE
			{
				(lex.p)++
				lex.cs = 0
				goto _out
			}
		}
		goto st0
	tr5:
//line lexer.rl:68
		lex.te = (lex.p)
		(lex.p)--
		{
			if mdDebug == 1 {
				fmt.Printf("%q", string(lex.data[lex.ts:lex.te]))
			}
			out.token = wordToken{rawToken(string(lex.data[lex.ts:lex.te]))}
			tok = TOKEN
			{
				(lex.p)++
				lex.cs = 0
				goto _out
			}
		}
		goto st0
	tr10:
//line lexer.rl:76
		lex.te = (lex.p)
		(lex.p)--
		{
			if mdDebug == 1 {
				fmt.Print("+")
			}
			out.token = whitespaceToken{rawToken(string(lex.data[lex.ts:lex.te]))}
			tok = WHITESPACE
			{
				(lex.p)++
				lex.cs = 0
				goto _out
			}
		}
		goto st0
	tr12:
//line lexer.rl:84
		lex.te = (lex.p) + 1
		{
			if mdDebug == 1 {
				fmt.Print("T_CODE_BLOCK")
			}
			out.token = fencedCodeblockToken{rawToken(string(lex.data[lex.ts:lex.te]))}
			tok = T_CODE_BLOCK
			{
				(lex.p)++
				lex.cs = 0
				goto _out
			}
		}
		goto st0
	st0:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof0
		}
	st_case_0:
//line NONE:1
		lex.ts = (lex.p)

//line lexer.go:167
		switch lex.data[(lex.p)] {
		case 9:
			goto st6
		case 10:
			goto tr2
		case 13:
			goto tr2
		case 32:
			goto st6
		case 96:
			goto st7
		case 126:
			goto st9
		}
		goto st1
	st1:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st2
		case 126:
			goto st4
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st3
		case 126:
			goto st4
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st3:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto tr5
		case 126:
			goto st4
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st4:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st2
		case 126:
			goto st5
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st5:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st2
		case 126:
			goto tr5
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st6:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		switch lex.data[(lex.p)] {
		case 9:
			goto st6
		case 32:
			goto st6
		}
		goto tr10
	st7:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st8
		case 126:
			goto st4
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st8:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto tr12
		case 126:
			goto st4
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st9:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st2
		case 126:
			goto st10
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st10:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		switch lex.data[(lex.p)] {
		case 13:
			goto tr5
		case 32:
			goto tr5
		case 96:
			goto st2
		case 126:
			goto tr12
		}
		if 9 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 10 {
			goto tr5
		}
		goto st1
	st_out:
	_test_eof0:
		lex.cs = 0
		goto _test_eof
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof
	_test_eof3:
		lex.cs = 3
		goto _test_eof
	_test_eof4:
		lex.cs = 4
		goto _test_eof
	_test_eof5:
		lex.cs = 5
		goto _test_eof
	_test_eof6:
		lex.cs = 6
		goto _test_eof
	_test_eof7:
		lex.cs = 7
		goto _test_eof
	_test_eof8:
		lex.cs = 8
		goto _test_eof
	_test_eof9:
		lex.cs = 9
		goto _test_eof
	_test_eof10:
		lex.cs = 10
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 1:
				goto tr5
			case 2:
				goto tr5
			case 3:
				goto tr5
			case 4:
				goto tr5
			case 5:
				goto tr5
			case 6:
				goto tr10
			case 7:
				goto tr5
			case 8:
				goto tr5
			case 9:
				goto tr5
			case 10:
				goto tr5
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:104

	return tok
}

func (lex *lexer) Error(e string) {
	fmt.Println("lexer error:", e)
}
