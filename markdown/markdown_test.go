package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParses(t *testing.T) {
	doc, err := Parse(`
foo  bat

	foo
		bar


~~~langname file name
	some
content

			content
~~~


`)
	assert.NoError(t, err)
	assert.Len(t, doc.CodeBlocks(), 1)
	codeBlock := doc.CodeBlocks()[0]
	assert.EqualValues(t, "langname", codeBlock.Language())
	assert.EqualValues(t, " file name", codeBlock.StringInfoAfterLanguageName())
	assert.EqualValues(t, `<blocks><block>+1"some"EOL</block><block>"content"EOL</block><block>EOL</block><block>+3"content"EOL</block></blocks>`, codeBlock.Content().String())
}
