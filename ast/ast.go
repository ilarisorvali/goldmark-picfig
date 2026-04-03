package ast

import (
	"fmt"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// KindFigure is a NodeKind of the Figure node.
var KindFigure = gast.NewNodeKind("Figure")

// A Figure struct represents a table of Markdown(GFM) text.
type FigurePicture struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *FigurePicture) Kind() gast.NodeKind {
	return KindFigure
}

// Dump implements Node.Dump
func (n *FigurePicture) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigure returns a new Table node.
func NewFigure() *FigurePicture {
	return &FigurePicture{}
}

// KindFigureImage is a NodeKind of the FigureImage node.
var KindFigureImage = gast.NewNodeKind("FigureImage")

// A FigureImage struct represents a table of Markdown(GFM) text.
type FigureImage struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *FigureImage) Kind() gast.NodeKind {
	return KindFigureImage
}

// Dump implements Node.Dump
func (n *FigureImage) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigurePicture returns a new FigurePicture node.
func NewFigureImage() *FigureImage {
	return &FigureImage{}
}

// KindFigureCaption is a NodeKind of the FigureCaption node.
var KindFigureCaption = gast.NewNodeKind("FigureCaption")

// A FigureCaption struct represents a node FigureCaption
type FigureCaption struct {
	gast.BaseBlock
}

// implements Node.Kind.
func (n *FigureCaption) Kind() gast.NodeKind {
	return KindFigureCaption
}

// implements Node.Dump
func (n *FigureCaption) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigureCaption returns a new FigureCaption node.
func NewFigureCaption() *FigureCaption {
	return &FigureCaption{}
}

// FigureHTMLRenderer is a renderer.NodeRenderer implementation that
// renders FigurePicture nodes.
type FigurePictureHTMLRenderer struct {
	renderImageLink bool
}

// NewFigureHTMLRenderer returns a new FigureHTMLRenderer.
func NewFigurePictureHTMLRenderer(renderImageLink bool) renderer.NodeRenderer {
	return &FigurePictureHTMLRenderer{renderImageLink: renderImageLink}
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *FigurePictureHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindFigure, r.renderFigure)
	reg.Register(KindFigureImage, r.renderFigureImage)
	reg.Register(KindFigureCaption, r.renderFigureCaption)
}

func (r *FigurePictureHTMLRenderer) renderFigure(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figure>\n")
	} else {
		_, _ = w.WriteString("</figure>\n")
	}
	return gast.WalkContinue, nil
}

func (r *FigurePictureHTMLRenderer) renderFigureImage(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if r.renderImageLink {
		if image, ok := n.FirstChild().(*gast.Image); ok {
			if entering {
				_, _ = w.WriteString(fmt.Sprintf("<a href=\"%s\">\n", string(image.Destination)))
			} else {
				_, _ = w.WriteString("</a>\n")
			}
		}
	}
	return gast.WalkContinue, nil
}

func (r *FigurePictureHTMLRenderer) renderFigureCaption(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figcaption><p>")
	} else {
		_, _ = w.WriteString("</p></figcaption>\n")
	}
	return gast.WalkContinue, nil
}
