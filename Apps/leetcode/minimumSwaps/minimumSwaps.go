package main

import (
    "fmt"
)
//4 3 1 2

func minimumSwaps(arr []int32) int32 {
  a := make(map[int32]int32)
  b := make(map[int32]int32)
  for i, x := range arr {
    a[int32(i)] = x
    b[x] = int32(i)
  } 
  swaps := 0
  for i, xA := range a{
    iter := i +1 
    if iter != xA {
      y := b[iter]
      a[y] = xA
      b[xA] = y

      swaps ++
    }else{
      continue
    }
  }

  return int32(swaps)
}
