package util

import (
	"encoding/json"
	"fmt"
)

func PrintJson(prefix string, obj any) {
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(prefix, string(jsonBytes))
}

func PrintStruct(prefix string, obj any) {
	fmt.Printf("%s %+v\n", prefix, obj)
}
