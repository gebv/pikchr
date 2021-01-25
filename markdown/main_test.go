package markdown

import (
	"os"
	"testing"

	"github.com/gebv/pikchr/markdown/syntax"
)

func TestMain(m *testing.M) {
	syntax.Debug()
	syntax.ErrorVerbose()

	os.Exit(m.Run())
}
