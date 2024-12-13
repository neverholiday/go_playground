package main

import "fmt"

type Students struct {
	Name      string `json:"name"`
	Age       int64  `json:"age"`
	ClassRoom string `json:"class"`
	Teacher   string `json:"teacher"`
}

type Response[T any] struct {
	Method  string `json:"method"`
	Payload T      `json:"payload"`
}

func main() {

	data := map[string]interface{}{
		"Number": 1,
		"Name":   "Nas",
	}

	fmt.Printf("number: %v (type: %T)\n",
		data["Number"],
		data["Number"])

	fmt.Printf("Name: %v (type: %T)\n",
		data["Name"],
		data["Name"])

}
