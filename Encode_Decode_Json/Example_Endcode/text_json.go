package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func marshalBasicTypes() {
	b, _ := json.Marshal(true)
	i, _ := json.Marshal(1)
	f, _ := json.Marshal(2.34)
	s, _ := json.Marshal("gopher")

	fmt.Printf("%s %s %s %s\n", b, i, f, s)
}

func marshalSliceAndMap() {
	fruits := []string{"apple", "peach", "pear"}
	fruitsJSON, _ := json.Marshal(fruits)

	fruitCounts := map[string]int{"apple": 5, "lettuce": 7}
	fruitCountsJSON, _ := json.Marshal(fruitCounts)

	fmt.Printf("%s %s\n", fruitsJSON, fruitCountsJSON)
}

func marshalCustomDataStructure() {
	data1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	data1JSON, _ := json.Marshal(data1)

	data2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	data2JSON, _ := json.Marshal(data2)

	fmt.Printf("%s %s\n", data1JSON, data2JSON)
}

func unmarshalJSONIntoCustomDataStructure() {
	jsonBytes := []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
	data := &response2{}
	json.Unmarshal(jsonBytes, data)

	fmt.Println(data)
}

func useEncoderToWriteJSONToStdout() {
	encoder := json.NewEncoder(os.Stdout)
	counts := map[string]int{"apple": 5, "lettuce": 7}
	encoder.Encode(counts)
}

func main() {
	marshalBasicTypes()
	marshalSliceAndMap()
	marshalCustomDataStructure()
	unmarshalJSONIntoCustomDataStructure()
	useEncoderToWriteJSONToStdout()
}
