package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonBlob := `{"timestamp":1420700998, "service":"daemon", "addr":"37/672", "ips":"121.14.241.43:CTL", "who":"pipm", "action":"Check", "msg":"daemonId:300313754 die res:direct die"}`
	var animals map[string]interface{}
	err := json.Unmarshal([]byte(jsonBlob), &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(animals)
}
