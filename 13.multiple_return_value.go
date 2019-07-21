package main

import "fmt"

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}

func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	return a
}

// 抛出多值
func getValues ()(int, int) {
  return 3, 7 
}

func main () {
  // 获取多值
  value1, value2 := getValues()
  fmt.Println(value1)
  fmt.Println(value2)

  // 使用_忽略不想要的值
  _, value3 := getValues()
  fmt.Println(value3)

  array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
  // array1 的值不会被改变
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println()

  slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
  // slice1 的值会被改变
	fmt.Printf("The modified slice: %v\n", slice2)
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println()

  // complexArray1 的值会被改变
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)
}
