package html

import (
	"encoding/json"
	"fmt"
)

func Client(id string, script string, data ...any) Element {
	tag := String(fmt.Sprintf("<!-- $$bundle:%s -->", script))

	if len(data) == 0 {
		return tag
	}

	json, err := json.Marshal(data[0])
	if err != nil {
		panic(err)
	}
	return Fragment(
		tag,
		Div(map[string]string{"data-id": id}),
		Tag(
			"script",
			map[string]string{
				"type":                     "application/json",
				"data-client-component-id": id,
			},
			String(string(json)),
		),
	)
}
