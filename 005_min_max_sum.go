package main

import (
	"fmt"
)

func changeType(arr []int32) []int64 {
	convertArr := make([]int64, len(arr))
	for i, v := range arr {
		convertArr[i] = int64(v)
	}
	return convertArr
}

func bubbleSort(arr []int32)[]int64{
	convertArr := changeType(arr)
	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(convertArr) - 1 {
			if convertArr[i] > convertArr[i + 1] {
				convertArr[i], convertArr[i + 1] = convertArr[i + 1], convertArr[i]
				isDone = false
			}
			i++
		}
	}
	return convertArr
}

func recursiveBubbleSort(arr []int32, size int) []int64 {
	if size == 1 {
			return changeType(arr)
	}

	var i = 0
	for i < size-1 {
			if arr[i] > arr[i+1] {
					arr[i], arr[i+1] = arr[i+1], arr[i]
			}

			i++
	}

	recursiveBubbleSort(arr, size - 1)
	return changeType(arr)
}

func insertionSort(arr []int32) []int64 {
	convertArr := changeType(arr)
	var i = 1
	for i < len(convertArr) {
		var j = i
		for j >= 1 && convertArr[j] < convertArr[j - 1] {
			convertArr[j], convertArr[j - 1] = convertArr[j - 1], convertArr[j]
			j--
		}
		i++
	}
	return convertArr
}

func selectionSort(arr []int32) []int64{
	convertArr := changeType(arr)
	var i = 1
	for i < len(convertArr) - 1 {
		var j = i + 1
		var minIndex = i

		if j < len(convertArr) {
			if convertArr[j] < convertArr[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			var temp = convertArr[i]
			convertArr[i] = convertArr[minIndex]
			convertArr[minIndex] = temp
		}
		i++
	}
	return convertArr
} // deu errado

func merge(fp []int, sp []int) []int {
	var n = make([]int, len(fp)+len(sp))

	var fpIndex = 0
	var spIndex = 0

	var nIndex = 0

	for fpIndex < len(fp) && spIndex < len(sp) {
			if fp[fpIndex] < sp[spIndex] {
					n[nIndex] = fp[fpIndex]
					fpIndex++
			} else if sp[spIndex] < fp[fpIndex] {
					n[nIndex] = sp[spIndex]
					spIndex++
			}

			nIndex++
	}

	for fpIndex < len(fp) {
			n[nIndex] = fp[fpIndex]

			fpIndex++
			nIndex++
	}

	for spIndex < len(sp) {
			n[nIndex] = sp[spIndex]

			spIndex++
			nIndex++
	}

	return n
} // verificar

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
			return arr
	}

	var fp = mergeSort(arr[0 : len(arr)/2])
	var sp = mergeSort(arr[len(arr)/2:])

	return merge(fp, sp)
} // verificar

func countSort(arr []int) []int {
	var max = arr[0]

	var i = 1
	for i < len(arr) {
			if arr[i] > max {
					max = arr[i]
			}

			i++
	}

	var indices = make([]int, max + 1)

	var j = 0
	for j < len(arr) {
			indices[arr[j]]++

			j++
	}

	var k = 1
	for k < len(indices) {
			indices[k] += indices[k - 1]

			k++
	}

	var result = make([]int, len(arr))

	var m = 0
	for m < len(arr) {
			result[indices[arr[m]] - 1] = arr[m]
			indices[arr[m]]--

			m++
	}

	return result
}

func miniMaxSum(arr []int32) {
	var min, max, sum int64
	// sortedArr := bubbleSort(arr)
	// sortedArr := recursiveBubbleSort(arr, len(arr))
	// sortedArr := insertionSort(arr)
	sortedArr := selectionSort(arr)
	for _, v := range sortedArr {
		sum += v
	}
	min = sum - sortedArr[len(sortedArr)-1]
	max = sum - sortedArr[0]
	fmt.Println(min, max)
}

func MiniMaxSum() {
	arr := []int32{623958417, 938071625, 714532089, 256741038, 467905213}
	miniMaxSum(arr)
}