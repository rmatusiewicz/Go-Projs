package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	min := int(arr[0])
	max := -1
	sum := 0

	for i := 0; i < 5; i++ {
		// fmt.Println(midSum)
		num := int(arr[i])
		sum = sum + num
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}

	fmt.Printf("%d %d", sum-max, sum-min)
	//7 69 2 221 8974
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// arrTemp := strings.Split(readLine(reader), " ")

	// var arr []int32

	// for i := 0; i < 5; i++ {
	// 	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	// 	checkError(err)
	// 	arrItem := int32(arrItemTemp)
	// 	arr = append(arr, arrItem)
	// }
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
