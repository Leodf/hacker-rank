package main

import (
	"fmt"
)

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
    var sum int64
    for _, v := range ar {
        sum += v
    }
    return sum
}


func VeryBigSum() {
    ar := []int64{1, 2, 3, 4, 5}
    sum := aVeryBigSum(ar)
    fmt.Println("VeryBigSum", sum)
}

