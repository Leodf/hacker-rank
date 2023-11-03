package main

import "fmt"

func plusMinus(arr []int32) {
	var positive, negative, zero float32
	var lengthArray = float32(len(arr))
	for _, v := range arr {
		switch {
		case v > 0: 
			positive++
		case v < 0:
			negative++
		default:
			zero++
		}
	}
	fmt.Println(positive/lengthArray)
	fmt.Println(negative/lengthArray)
	fmt.Println(zero/lengthArray)
}

func PlusMinus(){
	arr := []int32{1, 1, 0, -1, -1}
	fmt.Println("PlusMinus")
	plusMinus(arr)
}