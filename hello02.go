package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x,n); v<lim {
		return v
	} else {
		fmt.Println("%g >= %g", v, lim)
	}
	return lim
}

func main() {

	defer fmt.Println("World1")
	defer fmt.Println("World2")

	sum := 0
	for i:= 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum  < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	if sum < 10000{
		fmt.Println("sum < 10000")
	}

	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)

	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s.", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}

	fmt.Println("Hello")

	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}

}
