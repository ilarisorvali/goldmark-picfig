package picfig

import (
	"github.com/yuin/goldmark"
	gparser "github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/ilarisorvali/goldmark-picfig/ast"
	myparser "github.com/ilarisorvali/goldmark-picfig/parser"
)

type extension struct {
	renderImageLink bool
	skipNoCaption   bool
}

var PicFig = &extension{}

func (f *extension) WithImageLink() *extension {
	f.renderImageLink = true
	return f
}

func (f *extension) WithSkipNoCaption() *extension {
	f.skipNoCaption = true
	return f
}

func (e *extension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		gparser.WithParagraphTransformers(
			util.Prioritized(myparser.NewFigureParagraphTransformer(e.skipNoCaption), 0),
		),
	)

	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(ast.NewFigurePictureHTMLRenderer(e.renderImageLink), 0),
	))
}
