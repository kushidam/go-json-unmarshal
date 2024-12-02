package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	bytes, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}

	var result map[string][]map[string]interface{}
	json.Unmarshal(bytes, &result)
	for key, list := range result {
		for _, item := range list {
			fmt.Printf("%s,%s,%s\n",key, item["name"], item["value"])
		}
	}
}
