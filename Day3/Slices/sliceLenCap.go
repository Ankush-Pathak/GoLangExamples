package main

import "fmt"

//What happends when slice capacity is exceeded

func main() {

	slice := make([]int, 2, 4)

	slice[0] = 100
	slice[1] = 200
	//slice[2] = 300 will cause runtime error

	fmt.Println("Len:", len(slice), "Cap:", cap(slice)) //Len: 2 Cap: 4

	slice = append(slice, 300)

	fmt.Println("Len:", len(slice), "Cap:", cap(slice)) //Len: 3 Cap: 4

	slice = append(slice, 400, 500) //Re-allocation to a new array with greater capacity

	fmt.Println("Len:", len(slice), "Cap:", cap(slice)) //Len: 5 Cap: 8

	slice2 := append(slice, 3000)

	slice2[0] = -1

	fmt.Println(&slice[0]) //Same as below
	fmt.Println(&slice2[0])

	//But slice2 will move on to different location once capacity exceeds

	slice2 = append(slice2, 1, 2, 3, 4, 5)

	fmt.Println(&slice[0]) //Different now
	fmt.Println(&slice2[0])
}
