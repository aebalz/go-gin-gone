package utils

import (
	"encoding/json"
	"fmt"
)

func PPrint(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(bytes))
}
