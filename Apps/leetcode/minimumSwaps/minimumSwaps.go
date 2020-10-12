package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
//4 3 1 2

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
  swaps := 0
  solved :=0

  this := [2]int32{-1, -1}
  withThis := [2]int32{-1, -1}

  for solved + 1 < len(arr) {
    zCnt := 0
    for i,num := range arr {
      fmt.Println(i, num)
      if num - int32(i+1) == 0{
        zCnt ++
        if zCnt == len(arr){
          return int32(swaps)
        }
        continue
      }
      if this[0] == -1 && num - int32(i+1) > 0 {
        fmt.Println("Here 1")
        this[0] = int32(i)
        this[1] = num
        continue
      }

      fmt.Println(num - int32(i+1))      
      if withThis[0] == -1 && num - int32(i+1) < 0 {
        fmt.Println("Here 2")
        withThis[0] = int32(i)
        withThis[1] = num
      }
      if this[0] != -1 && withThis[0] != -1 {
        fmt.Printf("\narr before %v\n", arr)
        solved = solved + swap(arr, this, withThis)
        fmt.Printf("\narr after %v\n", arr)
        fmt.Printf("\nSolved %d\n", solved)
        swaps ++
        fmt.Printf("\nSwaps %d\n", swaps)

        //Reset
        this = [2]int32{-1, -1}
        withThis = [2]int32{-1, -1}
        break
      }
    
    }
  }

  return int32(swaps)
}

func swap(ar []int32, a, b [2]int32) int{
  fmt.Printf("SWAP %v, %v", a ,b)
  ret :=0
  ar[a[0]] = b[1]
  ar[b[0]] = a[1]

  if ar[a[0]] - (a[0]+1) == 0 {
    ret ++
  }

  if ar[b[0]] - (b[0]+1) == 0{
    ret ++
  }
  // fmt.Println(ret)

  return ret

