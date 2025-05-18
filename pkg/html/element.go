package html

import (
	"fmt"

	"strings"
)

type Element interface {
	Render(...*Stream) string
	AddChild(child Element) Element
	AddAttribute(key, value string) Element
	AddText(text string) Element
}

type Node struct {
	Tag        string
	Attributes map[string]string
	Children   []Element
	Text       string
}

func NewNode(tag string) Element {
	return &Node{
		Tag:        tag,
		Attributes: make(map[string]string),
		Children:   make([]Element, 0),
		Text:       "",
	}
}

func (e *Node) Render(stream ...*Stream) string {
	if e.Text != "" {
		return e.Text
	}

	if e.Tag == "fragment" {
		return strings.Join(Map(e.Children, func(child Element) string {
			return child.Render(stream...)
		}), "")
	}

	children := strings.Join(Map(e.Children, func(child Element) string {
		return child.Render(stream...)
	}), "")

	attributes := strings.Join(Map(Entries(e.Attributes), func(entry Entry[string]) string {
		return fmt.Sprintf(`%s="%s"`, entry.Key, entry.Value)
	}), "")

	return fmt.Sprintf("<%s %s>%s</%s>", e.Tag, attributes, children, e.Tag)
}

func (e *Node) AddChild(child Element) Element {
	e.Children = append(e.Children, child)
	return e
}

func (e *Node) AddAttribute(key, value string) Element {
	e.Attributes[key] = value
	return e
}

func (e *Node) AddText(text string) Element {
	e.Text = text
	return e
}

func Tag(tag string, attributes map[string]string, children ...Element) Element {
	e := NewNode(tag)
	for key, value := range attributes {
		e.AddAttribute(key, value)
	}
	for _, child := range children {
		e.AddChild(child)
	}
	return e
}

func Value(value string) Element {
	return &Node{
		Tag:        "",
		Attributes: make(map[string]string),
		Children:   make([]Element, 0),
		Text:       value,
	}
}
