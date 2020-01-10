package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonData, _ := readJson("./a.json")

	var unknown interface{}
	_ = json.Unmarshal(jsonData, &unknown)

	//assert, switch
	m := unknown.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "type: stirng\nvalue: ", vv)
			fmt.Println("...............")
		case float64:
			fmt.Println(k, "type: float64\nvalue: ", vv)
			fmt.Println(".........")
		case map[string]interface{}:
			fmt.Println(k, "type: map[string]interface{}\nvalue: ", vv)
			for i, j := range vv {
				fmt.Println(i, ":", j)
			}
			fmt.Println("...............")
		case []interface{}:
			fmt.Println(k, "type: []interface{}\nvalue: ", vv)
			for i, j := range vv {
				fmt.Println(i, ":", j)
			}
			fmt.Println("...............")
		case bool:
			fmt.Println(k, "type: bool\nvalue: ", vv)
			fmt.Println("...............")
		default:

		}
	}
}

func readJson(fileName string) ([]byte, error) {
	fh, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer fh.Close()

	data, err := ioutil.ReadAll(fh)
	if err != nil {
		return nil, err
	}

	return data, nil
}
