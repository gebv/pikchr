package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// NOTE: content always end with EOL

	emptyContent := "<blocks></blocks>"
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "emtpy", want: emptyContent},
		// TODO: why error?
		// {name: "oneline", in: "foobar", want: `<blocks><block>"foobar"</block></blocks>`},
		{name: "oneline", in: "foobar\n", want: `<blocks><block>"foobar"EOL</block></blocks>`},
		{name: "", in: "foo bar\n", want: `<blocks><block>"foo"+1"bar"EOL</block></blocks>`},
		{name: "", in: "foo  bar\n", want: `<blocks><block>"foo"+2"bar"EOL</block></blocks>`},
		{name: "", in: "foo \tbar\n", want: `<blocks><block>"foo"+2"bar"EOL</block></blocks>`},
		{name: "", in: " \tbar\n", want: `<blocks><block>+2"bar"EOL</block></blocks>`},

		{name: "should be EOL between cdmarkers", in: "``````\n", want: emptyContent, wantErr: true},
		{name: "have start codeblock but without end marker", in: "```\n", want: emptyContent, wantErr: true},
		{name: "should be enter after string info", in: "```foobar```\n", want: emptyContent, wantErr: true},

		{name: "codeblock_empty", in: "```\n```\n", want: `<blocks><fenced_code_block><StringInfo><block>EOL</block></StringInfo><Content>/<Content></fenced_code_block></blocks>`},
		{name: "codeblock_with_stringinfo_emptycontent", in: "```foo bar\n```\n", want: `<blocks><fenced_code_block><StringInfo><block>"foo"+1"bar"EOL</block></StringInfo><Content>/<Content></fenced_code_block></blocks>`},
		{name: "codeblock-ok1", in: "```foo bar\nsome text\nmany\n line\n  line\n\n\n\n```\n", want: `<blocks><fenced_code_block><StringInfo><block>"foo"+1"bar"EOL</block></StringInfo><Content><block>"some"+1"text"EOL</block><block>"many"EOL</block><block>+1"line"EOL</block><block>+2"line"EOL</block><block>EOL</block><block>EOL</block><block>EOL</block>/<Content></fenced_code_block></blocks>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse([]byte(tt.in))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, got.String())
		})
	}
}

func TestRawCodeBlock(t *testing.T) {
	stringinfo := "foo bar\n"
	content := "some text\nmany\n line\n  line\n\n\n\n"
	got, err := Parse([]byte("```" + stringinfo + content + "```\n"))
	assert.NoError(t, err)
	assert.EqualValues(t, FENCEDCODEBLOCK, got.Content()[0].Kind())
	assert.EqualValues(t, stringinfo, got.Content()[0].(*CodeBlock).StringInfo().(Rawer).Raw())
	assert.EqualValues(t, content, got.Content()[0].(*CodeBlock).Content().Raw())
}
