package main

import "os"
import "fmt"
import "encoding/json"

type response1 struct {
  Page int
  Fruits []string
}

type response2 struct {
  Page int `json:"page"`
  Fruits []string `json:"fruits"`
}

func main () {
  // JSON序列化布尔值
  bolB, _ := json.Marshal(true)
  fmt.Println(string(bolB))

  intB, _ := json.Marshal(1)
  fmt.Println(string(intB))

  floatB, _ := json.Marshal(1.23)
  fmt.Println(string(floatB))

  stringB, _ := json.Marshal("string")
  fmt.Println(string(stringB))

  stringArray := []string{"a", "b", "c"}
  stringArrayB, _ := json.Marshal(stringArray)
  fmt.Println(string(stringArrayB))

  maps := map[string]int{"a": 1, "b": 2}
  mapB, _ := json.Marshal(maps)
  fmt.Println(string(mapB))

  response1D := &response1{
    Page: 1,
    Fruits: []string{"a", "b"}}
  response1B, _ := json.Marshal(response1D)
  fmt.Println(string(response1B))

  response2D := &response2{
    Page: 1,
    Fruits: []string{"a", "b"}}
  response2B, _ := json.Marshal(response2D)
  fmt.Println(string(response2B))

  var dat map[string]interface{}
  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

  if err := json.Unmarshal(byt, &dat); err != nil {
    panic(err)
  }
  fmt.Println(dat)
  
  strs := dat["strs"].([]interface{})
  fmt.Println("strasdfsa:", dat)
  fmt.Println("strasdfsa:", strs)
  fmt.Println("strasdfsa:", strs[0])
  str1 := strs[0].(string)
  fmt.Println(str1)

  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res := response2{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println(res)
  fmt.Println(res.Fruits[0])

  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple": 5, "lettuce": 7}
  enc.Encode(d)
}
