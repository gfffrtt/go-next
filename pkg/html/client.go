package html

import (
	"encoding/json"
)

func Client(id string, script string, data ...any) Element {
	if len(data) == 0 {
		return Div(map[string]string{"data-id": id})
	}

	json, err := json.Marshal(data[0])
	if err != nil {
		panic(err)
	}
	return Fragment(
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
