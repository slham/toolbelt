package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {
	doc, err := ReadFile("toolbelt", "doc.go")
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, `// The toolbelt package is the common library for work.
// It currently has the following sub-modules:
// * l
//   * A tool for standardizing logs into a JSON format
// * constants
//   * A tool for supplying common string; variables, formats, regular expressions, etc.
package toolbelt
`, string(doc))
}

func TestReadFileNoFile(t *testing.T) {
	doc, err := ReadFile("toolbelt", "shank")
	if err == nil{
		t.Error("should not have been able to read a file")
	}else{
		assert.Equal(t, "open /Users/sheldonhampton/dank/toolbelt/shank: no such file or directory", err.Error())
		assert.Equal(t, []byte(nil), doc)
	}
}
