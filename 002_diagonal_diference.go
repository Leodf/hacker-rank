package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/diagonal-difference/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

func diagonalDifference(arr [][]int32) int32 {
	var primaryDiagonal int32
	var secondaryDiagonal int32
	for index, row := range arr {
		lastIndex := len(row)-1
		primaryDiagonal += row[index]
		secondaryDiagonal += row[lastIndex-index]
	}
	value := primaryDiagonal - secondaryDiagonal
	if value < 0 {
		value *= -1
		return value
	}
	return value
}

func DiagonalDifference() {
	arr := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	result := diagonalDifference(arr)
	fmt.Println("DiagonalDifference", result)
}