package main

import (
	"fmt"
	"math"
	"time"
)

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64
type F float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(-f)
}

func Scale1 (v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2 (v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func (f F) M(){
	fmt.Println(f)
}

func describe(i I){
	fmt.Printf("(%v, %T) ", i, i)
}

func do(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main(){
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	(&v).Scale(10)
	fmt.Println(v.Abs())

	Scale1(&v, 10)
	fmt.Println(Abs2(v))

	var a Abser
	a = f
	fmt.Println(a.Abs())

	var b = &v
	fmt.Println(b.Abs())

	//var i T = T {"hello"}
	//i.M()
	////////////////////////////////////////////

	fmt.Println("\n")

	var i I
	i = &T{"Hello"}
	describe(i)
	//i.M()

	i = F(math.Pi)
	describe(i)
	//i.M()

	var j I
	describe(j)
	//j.M()

	fmt.Println("\n")

	var i1 interface{} = "hello"
	s := i1.(string)
	fmt.Println(s)

	s, ok := i1.(string)
	fmt.Println(s, ok)

	//f, ok1 := i1.(float64)
	//fmt.Println(f, ok1)

	do(21)
	do("Hello")
	do(true)

	a1 := Person{"AAA", 12}
	z1 := Person{"BBB", 13}
	fmt.Println(a1, z1)

	if err := run(); err != nil{
		fmt.Println(err)
	}
}
