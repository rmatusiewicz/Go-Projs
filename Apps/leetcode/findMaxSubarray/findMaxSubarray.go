package main
import "fmt"

func main(){
  //
//
//
//      []         []
//      []         []
//      []  []     []
//  []  []  [] []  []
//----------------------------------------------------------------------
//[]  []  []     []
//[]  []         []
//    []         []
//               []
//               []
//
//
  //[-2,1,-3,4,-1,2,1,-5,4]
  //x := [5]int{10, 20, 30, 40, 50} 
  fmt.Println("TEST 1")
  input1 := []int{-2,1,-3,4,-1,2,1,-5,4}
  res1 := maxSubArray(input1)
  if res1 != 6 {
    fmt.Printf("Error: Expected 6 got %d, Input: %v", res1, input1)
  }
  
    fmt.Println("TEST 2")
  input2 := []int{1}
  res2 := maxSubArray(input2)
  if res2 != 1 {
    fmt.Printf("Error: Expected 1 got %d, Input: %v", res2, input2)
  }
  
    fmt.Println("TEST 3")
  input3 := []int{0}
  res3:= maxSubArray(input3)
  if res3!= 0 {
    fmt.Printf("Error: Expected 0 got %d, Input: %v", res3, input3)
  }
  
    
      fmt.Println("TEST 4")
  input4 := []int{-10, -13}
  res4:= maxSubArray(input4)
  if res4!= -10 {
    fmt.Printf("Error: Expected -10 got %d, Input: %v", res4, input4)
  }
}

//Notes
//* Negative num will never start or end the max sub sequence (unless only one element)
func maxSubArray(nums []int) int {
    if len(nums) == 1 {
      return nums[0]
    }
    
    // mini starts from negative to positive change
    // mini := 0
    
    // maxi starts from positive to negative change
    // maxi := 0
    var max int
    sum := nums[0]
    
    for i := 0; i <len(nums); i++ {
      fmt.Println(i, sum)
      // sum = sum + nums[i]
      if nums[i] < 0 {
        
      } else {
        //Add to localMax
      }
      
      
      if sum < 0 {
        max = sum - nums[i]
        sum = nil
      }
      
      // current := nums[i]
      
      //If >= 0 include in potential max
      
    }
    
    return max
}
