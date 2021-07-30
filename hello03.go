package main

import (
	"fmt"
	"strings"
	"math"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	X, Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func printSlice(s[]int){
	fmt.Printf("len=%d cap=%d %v ", len(s), cap(s), s)
}

type Vertex3 struct {
	Lat, Long float64
}

var m1 map[string]Vertex3

func compute(fn func(float64, float64) float64) float64{
	return fn(3, 4)
}

func adder() func(int) int{
	sum := 0
	return func (x int) int {
		sum += x
		return sum
	}
}

func main() {
	i, j := 42, 2700
	p := &i
	fmt.Println(*p)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(Vertex{1,2})

	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	q := &v
	q.X = 1e9
	fmt.Println(v.X)

	fmt.Println(v1, v2, v3, p)

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	var s[]int = primes[1:4]
	fmt.Println(s)

	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1,b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	s1 := []struct {
		i int
		b bool
	}{
		{1,true},
		{3,false},
		{5,true},
	}
	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	s2 = s2[1:4]
	fmt.Println(s2)
	fmt.Println(s2[:2])
	fmt.Println(s2[0:])

	s3 := []int{2,3,5,7,11,13}
	printSlice(s3)

	s3 = s3[:0]
	printSlice(s3)

	s3 = s3[:4]
	printSlice(s3)

	s3 = s3[2:]
	printSlice(s3)

	var s4[]int
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil{
		fmt.Println("nil")
	}

	board := [][]string{
		[]string{"1","2","3"},
		[]string{"4","5","6"},
		[]string{"7","8","9"},
	}

	for i:=0; i<len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}

	var s5 []int
	printSlice(s5)

	s5 = append(s5, 0)
	printSlice(s5)

	s5 = append(s5, 1)
	printSlice(s5)

	s5 = append(s5, 2, 3, 4)
	printSlice(s5)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	m1 = make(map[string]Vertex3)
	m1["Bell Labs"] = Vertex3{
		40.32, 78.32,
	}
	fmt.Println(m1["Bell Labs"])

	var m2 = map[string]Vertex3{
		"Bell Labs": Vertex3{
			49.1, 65.1,
		},
		"Google": Vertex3{
			32.1, 32.1,
		},
	}
	fmt.Println(m2)

	m3 := make(map[string]int)
	m3["answer"] = 42
	fmt.Println(m3["answer"])

	delete(m3, "answer")
	fmt.Println(m3["answer"])

	v1, ok := m3["answer"]
	fmt.Println("The value:", v1, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i:=0; i<5; i++ {
		fmt.Println(pos(i),neg(2*i+1))
	}
}
