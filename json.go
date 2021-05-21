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
func main() {
	bolB, _ := json.Marshal(true)
	intB, _ := json.Marshal(1)
	fltB, _ := json.Marshal(2.34)
	strB, _ := json.Marshal("gopher")
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	num := dat["num"].(float64)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	fmt.Println(string(bolB))
	fmt.Println(string(intB))
	fmt.Println(string(fltB))
	fmt.Println(string(strB))
	fmt.Println(string(slcB))
	fmt.Println(string(mapB))
	fmt.Println(string(res1B))
	fmt.Println(string(res2B))
	fmt.Println(dat)
	fmt.Println(num)
	fmt.Println(str1)
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}