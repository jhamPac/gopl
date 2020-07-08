package main

import "encoding/xml"

// Node represents arbitrary XML elements
type Node interface{}

// CharData represents data
type CharData string

// Element represents actual XML elements
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}
