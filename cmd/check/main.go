package main

import (
	"encoding/json"
	"os"
)

func main() {
	response := make([]string, 0)

	err := json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		panic(err)
	}
}
