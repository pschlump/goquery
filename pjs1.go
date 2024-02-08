package goquery

import "golang.org/x/net/html"

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

/*
CLI - Interpreter

>> read  file.html
>> get   url
>> .Find Pat
	-> print out selection found

>> quit,:q,:wq,\q etc.

*/
