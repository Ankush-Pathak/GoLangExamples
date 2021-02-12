package main

import "fmt"

func main() {
	map1 := map[string]int{"key": 100, "key2": 200} //string is key, int value

	fmt.Println(map1["key"])

	for key, val := range map1 {
		fmt.Println(key, val)
	}

	map1["key3"] = 400 //Update or add

	//"making" a map

	var map2 = make(map[int]string)
	map2[1] = "A"

	//delete
	delete(map2, 1)

	fmt.Println(map2)
}
