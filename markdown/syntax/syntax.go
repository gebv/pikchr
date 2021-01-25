package syntax

import "errors"

// Debug activates debug output to stdout.
func Debug() {
	mdDebug = 1
}

// ErrorVerbose adds detailed information about the parser error.
func ErrorVerbose() {
	mdErrorVerbose = true
}

// Parse parse input text as markdown text.
//
// Returns blocks of parsed text.
// NOTE: now supported
// - Fenced code blocks. String info and content.
// - Words block (paragraph with words, without parsing formatted text and other). At the end of EOL.
func Parse(in []byte) (Blocks, error) {
	lex := newLexer(in)
	e := mdNewParser().Parse(lex)
	if e != 0 {
		// TODO: more details
		return nil, errors.New("invalid format")
	}
	return lex.result, nil
}
