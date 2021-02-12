package main

//Suffix _test required to file

import "testing"

func TestSum(t *testing.T) {
	res := GetData("C")

	if res == 100 {
		t.Logf("Success")
	}
}
