package goquery

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

/*
type Document struct {
	*Selection
	Url      *url.URL
	rootNode *html.Node
}
*/

func (d *Document) GetRootNode() *html.Node {
	return d.rootNode
}

/*
type Selection struct {
	Nodes    []*html.Node
	document *Document
	prevSel  *Selection
}
*/

func (s *Selection) GetRootNode() *html.Node {
	return s.document.rootNode
}

// Render renders the HTML of the first item in the selection and writes it to
// the writer. It behaves the same as OuterHtml but writes to w instead of
// returning the string.
func RenderAll(w io.Writer, s *Selection) error {
	if s.Length() == 0 {
		return nil
	}
	// 	n := s.Get(0)
	for _, n := range s.Nodes {
		err := html.Render(w, n)
		if err != nil {
			return err
		}
		w.Write([]byte("\n"))
	}
	return nil
}

// OuterHtml returns the outer HTML rendering of the first item in
// the selection - that is, the HTML including the first element's
// tag and attributes.
//
// Unlike Html, this is a function and not a method on the Selection,
// because this is not a jQuery method (in javascript-land, this is
// a property provided by the DOM).
func OuterHtmlAll(s *Selection) (string, error) {
	var buf bytes.Buffer
	if err := RenderAll(&buf, s); err != nil {
		return "", err
	}
	return buf.String(), nil
}
