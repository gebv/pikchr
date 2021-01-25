package syntax

import (
	"fmt"
)

const (
	EOL                    = "eol"
	FENCEDCODEBLOCK        = "fenced_code_block"
	FENCEDCODEBLOCK_MARKER = "fenced_code_block_marker"
	SPACE                  = "whitespace"
	WORD                   = "word"
	BLOCK                  = "block"
)

// FencedCodeBlock it is fenced code block.
type FencedCodeBlock interface {
	Block
	StringInfo() Block
}

var _ FencedCodeBlock = (*CodeBlock)(nil)

// Block it is container with nested blocks.
type Block interface {
	Stringer
	// Kind returns kind of block.
	Kind() string
	// AddBlock adds block into current block.
	AddBlock(b ...Block)
	// Content returns all blocks.
	Content() Blocks
}

type Rawer interface {
	// Raw returns raw content of block.
	Raw() string
}

var _ Rawer = (Blocks)(nil)
var _ Rawer = (*LineBlock)(nil)

var _ Rawer = (*wordToken)(nil)
var _ Rawer = (*eolToken)(nil)
var _ Rawer = (*whitespaceToken)(nil)
var _ Rawer = (*fencedCodeblockToken)(nil)

type Stringer interface {
	// String returns as text of the content. Use for debuging or preview raw content.
	String() string
}

var _ Block = (*CodeBlock)(nil)
var _ Block = (*Blocks)(nil)
var _ Block = (*LineBlock)(nil)

// Token elementary syntax unit.
type Token interface {
	Stringer
	// Kind returns kind of token.
	Kind() string
}

var _ Token = (*wordToken)(nil)
var _ Token = (*eolToken)(nil)
var _ Token = (*whitespaceToken)(nil)
var _ Token = (*fencedCodeblockToken)(nil)

// NewCodeBlock returns code block.
func NewCodeBlock(start, end Token, b ...Block) *CodeBlock {
	return &CodeBlock{
		StartToken: start,
		EndToken:   end,
		Blocks:     b,
	}
}

// CodeBlock it is fenced code block. Implements Block interface.
type CodeBlock struct {
	Blocks               Blocks
	StartToken, EndToken Token
}

func (cb CodeBlock) StringInfo() Block {
	// NOTE: always has block (see parser rule)
	return cb.Blocks[0]
}

func (cb CodeBlock) Kind() string {
	return "fenced_code_block"
}

func (cb *CodeBlock) AddBlock(b ...Block) {
	cb.Blocks.AddBlock(b...)
}

func (cb CodeBlock) Content() Blocks {
	// NOTE: always has one block (see parser rule)
	return cb.Blocks[1:]
}

func (cb CodeBlock) String() string {
	var res string
	res += "<" + cb.Kind() + ">"
	res += "<StringInfo>"
	res += cb.StringInfo().String()
	res += "</StringInfo>"
	res += "<Content>"
	for _, b := range cb.Content() {
		res += b.String()
	}
	res += "/<Content>"
	res += "</" + cb.Kind() + ">"
	return res
}

// Blocks it is helper type. Used for store list of blocks. Implements Blocks interface.
type Blocks []Block

func (b Blocks) Kind() string {
	return "blocks"
}

func (bs *Blocks) AddBlock(b ...Block) {
	*bs = append(*bs, b...)
}

func (b Blocks) Content() Blocks {
	return b
}

func (bs Blocks) String() string {
	var res string
	res += "<" + bs.Kind() + ">"
	for _, b := range bs {
		res += b.String()
	}
	res += "</" + bs.Kind() + ">"
	return res
}

func (cb Blocks) Raw() string {
	var res string
	for _, b := range cb {
		if dat, ok := b.(Rawer); ok {
			res += dat.Raw()
		} else {
			res += "<BlockNotImplementRawer>"
		}

	}
	return res
}

// LineBlock it is helper type. Used for store list of tokens. Implements Block interface.
type LineBlock []Token

func (b LineBlock) Kind() string {
	return "block"
}

func (b LineBlock) Tokens() []Token {
	return b
}

func (b *LineBlock) AddToken(t ...Token) {
	*b = append(*b, t...)
}

func (b *LineBlock) AddBlock(...Block) {
	panic("not supported")
}

func (b *LineBlock) Content() Blocks {
	return Blocks{b}
}

func (cb LineBlock) Raw() string {
	var res string
	for _, b := range cb {
		if dat, ok := b.(Rawer); ok {
			res += dat.Raw()
		} else {
			res += "<TokenNotImplementRawer>"
		}
	}
	return res
}

func (b LineBlock) String() string {
	var res string
	res += "<" + b.Kind() + ">"
	for _, t := range b {
		res += t.String()
	}
	res += "</" + b.Kind() + ">"
	return res
}

func rawToken(in string) tokenContainer {
	return tokenContainer{
		raw: in,
	}
}

type tokenContainer struct {
	raw string
	// TODO: position on line
}

func (t tokenContainer) Raw() string {
	return t.raw
}

type wordToken struct {
	tokenContainer
}

func (t wordToken) Kind() string {
	return "word"
}

func (t wordToken) String() string {
	return fmt.Sprintf("%q", t.Raw())
}

type eolToken struct {
	tokenContainer
}

func (t eolToken) Kind() string {
	return "eol"
}

func (t eolToken) String() string {
	return "EOL"
}

type whitespaceToken struct {
	tokenContainer
}

func (t whitespaceToken) Kind() string {
	return "whitespace"
}

func (t whitespaceToken) String() string {
	return fmt.Sprintf("+%d", len(t.Raw()))
}

type fencedCodeblockToken struct {
	tokenContainer
}

func (t fencedCodeblockToken) Kind() string {
	return "fenced_code_block_marker"
}

func (t fencedCodeblockToken) String() string {
	return "<FencedCodeBlockMarker>"
}
