package picfig

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	"github.com/ilarisorvali/goldmark-picfig/ast"
)

type extension struct{}

var PicFig = &extension{}

func (e *extension) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(ast.NewFigureHTMLRenderer(e.renderImageLink), 0),
	))
}
