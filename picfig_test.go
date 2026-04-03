package picfig_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark"

	"github.com/ilarisorvali/goldmark-picfig"
)

func TestFigureRendering(t *testing.T) {
	md := goldmark.New(
		goldmark.WithExtensions(
			picfig.PicFig,
		),
	)

	input := `![alt text](assets/fig.jpg)
This is a caption`

	var buf bytes.Buffer
	if err := md.Convert([]byte(input), &buf); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}

	output := buf.String()

	// Basic checks
	if !strings.Contains(output, "<figure>") {
		t.Errorf("expected <figure> tag, got:\n%s", output)
	}

	if !strings.Contains(output, "<figcaption>") {
		t.Errorf("expected <figcaption>, got:\n%s", output)
	}

	if !strings.Contains(output, "This is a caption") {
		t.Errorf("caption missing, got:\n%s", output)
	}

	if !strings.Contains(output, "img") {
		t.Errorf("image not rendered, got:\n%s", output)
	}
}
