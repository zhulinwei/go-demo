package main

import "fmt"

func main () {
  map1 := make(map[string]string)
  map1["key1"] = "value1"
  map1["key2"] = "value2"

  fmt.Println("map1:", map1)

  value1 := map1["key1"]
  fmt.Println("value1", value1)
  fmt.Println("map1 length:", len(value1))

  delete(map1, "key1")
  fmt.Println("map1", map1)
  
  map2 := map[string]int{ "foo": 1, "bar": 2 }
  fmt.Println("map2", map2)
}
