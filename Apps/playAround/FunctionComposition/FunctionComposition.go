package main

import "fmt"
import "math"

func Compose(f, g func(x float64) float64) func(x float64) float64{
	return func(x float64) float64{
		return f(g(x))
	}
}

func main(){
	fmt.Println(Compose(math.Sin, math.Cos)(0.5))
}