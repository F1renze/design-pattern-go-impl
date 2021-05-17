package visitor

import "testing"

func TestVisitor(t *testing.T) {
	c := NewResourceContainer()

	c.Attach(NewPdfFile("pdf1"))
	c.Attach(NewWordFile("word1"))
	c.Attach(NewPPTFile("ppt1"))

	v1 := NewExtractor()
	v2 := NewCompressor()

	c.Accept(v1)
	c.Accept(v2)
}
