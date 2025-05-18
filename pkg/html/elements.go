package html

import (
	"strconv"
)

func String(value string) Element {
	return Value(value)
}

func Int[T ~int](value T) Element {
	return Value(strconv.Itoa(int(value)))
}

func Float[T ~float64](value T) Element {
	return Value(strconv.FormatFloat(float64(value), 'f', -1, 64))
}

func Bool(value bool) Element {
	return Value(strconv.FormatBool(value))
}

func Span(attributes map[string]string, children ...Element) Element {
	return Tag("span", attributes, children...)
}

func P(attributes map[string]string, children ...Element) Element {
	return Tag("p", attributes, children...)
}

func Div(attributes map[string]string, children ...Element) Element {
	return Tag("div", attributes, children...)
}

func A(attributes map[string]string, children ...Element) Element {
	return Tag("a", attributes, children...)
}

func Img(attributes map[string]string) Element {
	return Tag("img", attributes)
}

func Input(attributes map[string]string) Element {
	return Tag("input", attributes)
}

func Button(attributes map[string]string, children ...Element) Element {
	return Tag("button", attributes, children...)
}

func Form(attributes map[string]string, children ...Element) Element {
	return Tag("form", attributes, children...)
}

func Label(attributes map[string]string, children ...Element) Element {
	return Tag("label", attributes, children...)
}

func Select(attributes map[string]string, children ...Element) Element {
	return Tag("select", attributes, children...)
}

func Option(attributes map[string]string, children ...Element) Element {
	return Tag("option", attributes, children...)
}

func Textarea(attributes map[string]string, children ...Element) Element {
	return Tag("textarea", attributes, children...)
}

func Html(attributes map[string]string, children ...Element) Element {
	return Tag("html", attributes, children...)
}

func Head(attributes map[string]string, children ...Element) Element {
	return Tag("head", attributes, children...)
}

func Body(attributes map[string]string, children ...Element) Element {
	return Tag("body", attributes, children...)
}

func Header(attributes map[string]string, children ...Element) Element {
	return Tag("header", attributes, children...)
}

func Main(attributes map[string]string, children ...Element) Element {
	return Tag("main", attributes, children...)
}

func Footer(attributes map[string]string, children ...Element) Element {
	return Tag("footer", attributes, children...)
}

func Section(attributes map[string]string, children ...Element) Element {
	return Tag("section", attributes, children...)
}

func Article(attributes map[string]string, children ...Element) Element {
	return Tag("article", attributes, children...)
}

func Aside(attributes map[string]string, children ...Element) Element {
	return Tag("aside", attributes, children...)
}

func Nav(attributes map[string]string, children ...Element) Element {
	return Tag("nav", attributes, children...)
}

func Ol(attributes map[string]string, children ...Element) Element {
	return Tag("ol", attributes, children...)
}

func Ul(attributes map[string]string, children ...Element) Element {
	return Tag("ul", attributes, children...)
}

func Li(attributes map[string]string, children ...Element) Element {
	return Tag("li", attributes, children...)
}

func Table(attributes map[string]string, children ...Element) Element {
	return Tag("table", attributes, children...)
}

func Tr(attributes map[string]string, children ...Element) Element {
	return Tag("tr", attributes, children...)
}

func Td(attributes map[string]string, children ...Element) Element {
	return Tag("td", attributes, children...)
}

func Th(attributes map[string]string, children ...Element) Element {
	return Tag("th", attributes, children...)
}

func Tbody(attributes map[string]string, children ...Element) Element {
	return Tag("tbody", attributes, children...)
}

func Tfoot(attributes map[string]string, children ...Element) Element {
	return Tag("tfoot", attributes, children...)
}

func Thead(attributes map[string]string, children ...Element) Element {
	return Tag("thead", attributes, children...)
}

func Meta(attributes map[string]string) Element {
	return Tag("meta", attributes)
}

func Title(attributes map[string]string, children ...Element) Element {
	return Tag("title", attributes, children...)
}

func Link(attributes map[string]string) Element {
	return Tag("link", attributes)
}

func Template(attributes map[string]string, children ...Element) Element {
	return Tag("template", attributes, children...)
}

func Script(attributes map[string]string, children ...Element) Element {
	return Tag("script", attributes, children...)
}

func Fragment(children ...Element) Element {
	return Tag("fragment", map[string]string{}, children...)
}
