package markdown

import "github.com/gebv/pikchr/markdown/syntax"

func Parse(in string) (*Markdown, error) {
	blocks, err := syntax.Parse([]byte(in))
	if err != nil {
		return nil, err
	}
	return &Markdown{content: blocks}, nil
}

type Markdown struct {
	content syntax.Blocks
}

func (m *Markdown) CodeBlocks() []MarkdownCodeBlock {
	res := []MarkdownCodeBlock{}
	for _, b := range m.content.Content() {
		if b.Kind() == syntax.FENCEDCODEBLOCK {
			res = append(res, MarkdownCodeBlock{content: b.(*syntax.CodeBlock)})
		}
	}
	return res
}

type MarkdownCodeBlock struct {
	content *syntax.CodeBlock
}

func (b MarkdownCodeBlock) Language() string {
	// NOTE: shuld be one
	stringInfo := b.StringInfo().Content()[0]
	okStringInfo, ok := stringInfo.(*syntax.LineBlock)
	if !ok {
		return ""
	}
	if len(okStringInfo.Tokens()) == 0 {
		return ""
	}
	firstWord := okStringInfo.Tokens()[0]
	if firstWord.Kind() != syntax.WORD {
		return ""
	}
	return firstWord.(syntax.Rawer).Raw()
}

func (b MarkdownCodeBlock) StringInfoAfterLanguageName() string {
	// NOTE: shuld be one
	stringInfo := b.StringInfo().Content()[0]
	okStringInfo, ok := stringInfo.(*syntax.LineBlock)
	if !ok {
		return ""
	}
	if len(okStringInfo.Tokens()) == 0 {
		return ""
	}
	afterFirstWord := okStringInfo.Tokens()[1:]
	if len(afterFirstWord) <= 1 {
		return ""
	}
	var res string
	for _, word := range afterFirstWord {
		if word.Kind() == syntax.EOL {
			return res
		}

		res += word.(syntax.Rawer).Raw()
	}
	panic("should be EOL on end line")
}

func (b MarkdownCodeBlock) StringInfo() syntax.Block {
	return b.content.StringInfo()
}

func (b MarkdownCodeBlock) Content() syntax.Blocks {
	return b.content.Content()
}
