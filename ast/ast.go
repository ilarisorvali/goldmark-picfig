package ast

import (
	"fmt"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// KindFigure is a NodeKind of the Figure node.
var KindFigure = gast.NewNodeKind("PictureFigure")

// A Figure struct represents a table of Markdown(GFM) text.
type PictureFigure struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *PictureFigure) Kind() gast.NodeKind {
	return KindFigure
}

// Dump implements Node.Dump
func (n *PictureFigure) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewFigure returns a new Table node.
func NewFigure() *PictureFigure {
	return &PictureFigure{}
}

// KindFigureImage is a NodeKind of the FigureImage node.
var KindFigureImage = gast.NewNodeKind("FigureImage")

// A FigureImage struct represents a table of Markdown(GFM) text.
type Picture struct {
	gast.BaseBlock
}

// Kind implements Node.Kind.
func (n *Picture) Kind() gast.NodeKind {
	return KindFigureImage
}

// Dump implements Node.Dump
func (n *Picture) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, func(level int) {
	})
}

// NewPictureFigure returns a new PictureFigure node.
func NewFigureImage() *Picture {
	return &Picture{}
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
// renders PictureFigure nodes.
type PictureFigureHTMLRenderer struct {
	renderImageLink bool
}

// NewFigureHTMLRenderer returns a new FigureHTMLRenderer.
func NewPictureFigureHTMLRenderer(renderImageLink bool) renderer.NodeRenderer {
	return &PictureFigureHTMLRenderer{renderImageLink: renderImageLink}
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *PictureFigureHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindFigure, r.renderFigure)
	reg.Register(KindFigureImage, r.renderPictureFigure)
	reg.Register(KindFigureCaption, r.renderFigureCaption)
}

func (r *PictureFigureHTMLRenderer) renderFigure(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figure>\n")
	} else {
		_, _ = w.WriteString("</figure>\n")
	}
	return gast.WalkContinue, nil
}

func (r *PictureFigureHTMLRenderer) renderPictureFigure(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
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

func (r *PictureFigureHTMLRenderer) renderFigureCaption(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		_, _ = w.WriteString("<figcaption><p>")
	} else {
		_, _ = w.WriteString("</p></figcaption>\n")
	}
	return gast.WalkContinue, nil
}
