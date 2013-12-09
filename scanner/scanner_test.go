package scanner

import (
	"bytes"
	"testing"
)

func TestBold(t *testing.T) {
	inDoc := `some **bold text** we have.`
	expected := `some **bold text** we have.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestItalic(t *testing.T) {
	inDoc := `some //italic text// we have.`
	expected := `some *italic text* we have.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestUnderline(t *testing.T) {
	inDoc := `some __underline text__ we have.`
	expected := `some underline text we have.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestMonospace(t *testing.T) {
	inDoc := `some ''monospace text'' we have.`
	expected := `some ` + "`" + `monospace text` + "`" + ` we have.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	inDoc := `some <del>strikethrough text</del> we have.`
	expected := `some ~~strikethrough text~~ we have.`
	outDoc := ConvertDoc([]byte(inDoc))
	if !bytes.Equal(outDoc, []byte(expected)) {
		t.Fail()
	}
}

func TestNestedFormatting(t *testing.T) {
	inDoc := `**bold and // italic// and __und-erlined text
with ''monospace'' and <del>deleted </del> mixed__ and nested  .**`
	expected := `**bold and * italic* and und-erlined text
with ` + "`" + `monospace` + "`" + ` and ~~deleted ~~ mixed and nested  .**`
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
