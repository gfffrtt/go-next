package html

import (
	"fmt"
)

type SuspenseNode struct {
	Id       string
	Fallback Element
	Children func() Element
}

func NewSuspenseNode(id string, fallback Element, children func() Element) *SuspenseNode {
	return &SuspenseNode{Id: id, Fallback: fallback, Children: children}
}

func (e *SuspenseNode) Render(stream ...*Stream) string {
	go func() {
		stream[0].Add()
		children := Fragment(
			Template(map[string]string{" data-suspense-id": e.Id}, e.Children()),
			Script(map[string]string{}, String(fmt.Sprintf(`
(() => {
	const template = document.querySelector('template[data-suspense-id="%s"]');
	const fallback = document.querySelector('div[data-fallback-id="%s"]');
	fallback.replaceWith(template.content); 
})();
		`, e.Id, e.Id))),
		)
		stream[0].Write(children.Render(stream...))
	}()
	return e.Fallback.Render(stream...)
}

func (e *SuspenseNode) AddChild(child Element) Element {
	return e
}

func (e *SuspenseNode) AddAttribute(key, value string) Element {
	return e
}

func (e *SuspenseNode) AddText(text string) Element {
	return e
}

func Suspense(element func() Element, fallback ...Element) Element {
	id, err := Id()
	if err != nil {
		panic(err)
	}

	fallbackParent := Div(
		map[string]string{
			"data-fallback-id": id,
			"style":            "display: contents;",
		},
	)

	if len(fallback) == 0 {
		return NewSuspenseNode(id, fallbackParent, element)
	}

	return NewSuspenseNode(id, fallbackParent.AddChild(fallback[0]), element)
}
