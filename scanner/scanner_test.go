package scanner

import (
	"bytes"
	"testing"
)

func TestBasicFormatting(t *testing.T) {
	inDoc := `**bold and //italic// and __underlined text
with ''monospace'' mixed__ and nested**`
	expected := `**bold and *italic* and underlined text
with ` + "`" + `monospace` + "`" + ` mixed and nested**`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestHeading(t *testing.T) {
	inDoc := `====== heading 1 ======
===== h2 =====
==== h3 ====
=== h4 ===
== h5 with some //italic// text ==`
	expected := `# heading 1 
## h2 
### h3 
#### h4 
##### h5 with some *italic* text `
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestCode(t *testing.T) {
	inDoc := `<code>Some code that
should ignore **bold and //italic//
text**.</code>But should //work// afterwards.`
	expected := "```" + `Some code that
should ignore **bold and //italic//
text**.` + "```" + `But should *work* afterwards.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestFile(t *testing.T) {
	inDoc := `<file>Some code that
should ignore **bold and //italic//
text**.</file>But should //work// afterwards.`
	expected := "```" + `Some code that
should ignore **bold and //italic//
text**.` + "```" + `But should *work* afterwards.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}
