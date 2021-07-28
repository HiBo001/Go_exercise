package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
)

//var c, python, java bool
var i, j int = 1, 2
var (
	Tobe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5+12i)
)
const Pi = 3.14
const (
	Big = 1<<50
	Smalll = Big >>54
)



func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Println("Hello world")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favourite number is", rand.Intn(8))
	fmt.Println(math.Pi)
	fmt.Println(add(42,33))
	fmt.Println(add1(42,33))
	a, b := swap("tang", "haibo")
	fmt.Println(a, b)
	fmt.Println(split(10))
	//var i int
	//var c, python, java = true, false, "no!"
	k := 3
	c, python, java := true, false, "no!" // := only allowed in a function
	fmt.Println(k, c, python, java)
	fmt.Println("Type: %T Value: %v", Tobe, Tobe)
	fmt.Println("Type: %T Value: %v", MaxInt, MaxInt)
	fmt.Println("Type: %T Value: %v", z, z)
	u := 42.342
	v := uint(u)
	fmt.Println("v is %v", v)
	fmt.Println(Pi)
	const world = "world"
	const Truth = true
	fmt.Println(world)
	fmt.Println(Truth)
	fmt.Println(Big)
}

func add1(x, y int) int {
	return x + y
}
