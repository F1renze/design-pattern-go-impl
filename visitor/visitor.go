package visitor

import "fmt"

// IResourceFile 资源类
type IResourceFile interface {
	Accept(v IVisitor)
}

type BaseFile struct {
	filePath string
}

func NewPdfFile(filePath string) IResourceFile {
	return &PdfFile{BaseFile{filePath: filePath}}
}

type PdfFile struct {
	BaseFile
}

func (f *PdfFile) Accept(v IVisitor) {
	v.VisitPdf(f)
}

func NewWordFile(filePath string) IResourceFile {
	return &WordFile{BaseFile{filePath: filePath}}
}

type WordFile struct {
	BaseFile
}

func (f *WordFile) Accept(v IVisitor) {
	v.VisitWord(f)
}

func NewPPTFile(filePath string) IResourceFile {
	return &PPTFile{BaseFile{filePath: filePath}}
}

type PPTFile struct {
	BaseFile
}

func (f *PPTFile) Accept(v IVisitor) {
	v.VisitPPT(f)
}

// IVisitor 操作类
type IVisitor interface {
	VisitPPT(f *PPTFile)
	VisitWord(f *WordFile)
	VisitPdf(f *PdfFile)
}

func NewCompressor() IVisitor {
	return &Compressor{}
}

type Compressor struct {
}

func (c Compressor) VisitPPT(f *PPTFile) {
	fmt.Println("compress ppt")
}

func (c Compressor) VisitWord(f *WordFile) {
	fmt.Println("compress word")
}

func (c Compressor) VisitPdf(f *PdfFile) {
	fmt.Println("compress pdf")
}

func NewExtractor() IVisitor {
	return &Extractor{}
}

type Extractor struct {
}

func (e Extractor) VisitPPT(f *PPTFile) {
	fmt.Println("extract ppt")
}

func (e Extractor) VisitWord(f *WordFile) {
	fmt.Println("extract word")
}

func (e Extractor) VisitPdf(f *PdfFile) {
	fmt.Println("extract pdf")
}

type IContainer interface {
	Attach(f IResourceFile)
	Detach(f IResourceFile)
	Accept(v IVisitor)
}

func NewResourceContainer() IContainer {
	return &ResourceContainer{resources: map[IResourceFile]struct{}{}}
}

// ResourceContainer 聚合资源的容器
type ResourceContainer struct {
	resources map[IResourceFile]struct{}
}

func (r *ResourceContainer) Attach(f IResourceFile) {
	_, ok := r.resources[f]
	if !ok {
		r.resources[f] = struct{}{}
	}
}

func (r *ResourceContainer) Detach(f IResourceFile) {
	_, ok := r.resources[f]
	if ok {
		delete(r.resources, f)
	}
}

func (r *ResourceContainer) Accept(v IVisitor) {
	for r, _ := range r.resources {
		r.Accept(v)
	}
}
