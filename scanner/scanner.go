package scanner

import (
	"bytes"
)

type itemType int

const (
	itemText itemType = iota
	itemHeading1
	itemHeading2
	itemHeading3
	itemHeading4
	itemHeading5
	itemBold
	itemItalic
	itemUnderline
	itemMonospace
	itemStartLink
	itemEndLink
	itemStartPic
	itemEndPic
	itemStartCode
	itemEndCode
	itemStartFile
	itemEndFile
	itemEOF
)

type item struct {
	start      int
	end        int
	origString string
	replString string
	typ        itemType
}

type markdown struct {
	doc         bytes.Buffer
	curStack    map[int]itemType
	curStackPos int
}

func ConvertDoc(doc []byte) []byte {
	md := newMarkdown()
	c := make(chan item)
	go parseDoc(doc, c)
	for {
		item := <-c
		if item.typ == itemEOF {
			break
		}
		md.processItem(item)
	}
	return md.doc.Bytes()
}

func newMarkdown() *markdown {
	stack := make(map[int]itemType, 0)
	return &markdown{curStack: stack, curStackPos: 0}
}

func (md *markdown) processItem(i item) {
	switch i.typ {
	case itemText:
		md.doc.WriteString(i.origString)
	case itemHeading1, itemHeading2, itemHeading3, itemHeading4, itemHeading5:
		if !md.validContext(i) {
			md.doc.WriteString(i.origString)
			break
		}
		if md.checkStack(i) {
			md.doc.WriteString(i.replString)
		}
	case itemBold, itemItalic, itemUnderline, itemMonospace, itemStartCode, itemEndCode, itemStartFile, itemEndFile:
		if !md.validContext(i) {
			md.doc.WriteString(i.origString)
			break
		}
		md.checkStack(i)
		md.doc.WriteString(i.replString)
	}
}

func (md *markdown) validContext(i item) bool {
	if md.curStackPos == 0 {
		return true
	}
	// Simply ignore the obvious cases:
	// For <code> and <file> blocks, only allow the respective end block.
	// Allow all other cases.
	switch md.curStack[md.curStackPos-1] {
	case itemStartCode:
		if i.typ == itemEndCode {
			return true
		} else {
			return false
		}
	case itemStartFile:
		if i.typ == itemEndFile {
			return true
		} else {
			return false
		}
	}
	return true
}

// checkStack manages the stack of current items. This is required for types
// that have opening and closing tags in dokuwiki, but only opening tags in
// markdown. It is also useful for more complicated nested types.
// It returns true if the item marks an opening tag, and false otherwise.
func (md *markdown) checkStack(i item) bool {
	if md.curStackPos == 0 || md.curStack[md.curStackPos-1] != i.typ {
		md.curStack[md.curStackPos] = i.typ
		md.curStackPos++
		return true
	}
	delete(md.curStack, md.curStackPos)
	md.curStackPos--
	return false
}
