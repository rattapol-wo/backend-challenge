package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func PrintStructJson(targets ...interface{}) {
	var values []string
	for i, target := range targets {
		switch reflect.TypeOf(target).Kind() {
		case reflect.String:
			values = append(values, target.(string))
		default:
			fooByte, _ := json.MarshalIndent(&target, "", "\t")
			values = append(values, string(fooByte))
		}

		if i != len(targets) {
			fmt.Print(values[i])
			fmt.Println(` `)
		} else {
			fmt.Println(values[i])
		}
	}
}