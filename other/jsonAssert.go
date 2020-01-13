package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	MA := []byte(`{"name": "Massachusetts", "area": 27336, "water": 25.7, "senators":["John Kerry", "Scott Brown"]}`)
	var object interface{}
	err := json.Unmarshal(MA, &object)
	if err != nil {
		fmt.Println(err)
	} else {
		jsonObject := object.(map[string]interface{})
		fmt.Println(jsonObjectAsString(jsonObject))
	}
}

func jsonObjectAsString(jsonObject map[string]interface{}) string {
	var buffer bytes.Buffer
	comma := ""
	buffer.WriteString("{")
	for key, value := range jsonObject {
		buffer.WriteString(comma)
		switch value := value.(type) {
		case string:
			_, _ = fmt.Fprintf(&buffer, "%q: %q", key, value)
		case bool:
			_, _ = fmt.Fprintf(&buffer, "%q: %t", key, value)
		case float64:
			_, _ = fmt.Fprintf(&buffer, "%q: %f", key, value)
		case nil:
			_, _ = fmt.Fprintf(&buffer, "%q: nul", key)
		case []interface{}:
			_, _ = fmt.Fprintf(&buffer, "%q: [", key)
			innerComma := ""
			for _, value := range value {
				if s, ok := value.(string); ok {
					_, _ = fmt.Fprintf(&buffer, "%s%q", innerComma, s)
					innerComma = ","
				}
			}
			buffer.WriteString("]")
		}
		comma = ","
	}
	buffer.WriteString("}")
	return buffer.String()
}
