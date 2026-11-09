package main

import "fmt"

const x = 10
const y int32 = 15

const (
	c = 10
	d = 20
)

func integerdemo() {
	var smallnum uint8
	smallnum = 20
	fmt.Println(smallnum)
	var myint int = 2333
	fmt.Println(myint)
	var smallfloat float32 = 23.444444
	fmt.Println(smallfloat)
	var mybigfloat float64 = 346.6666666667653
	var mycomplex complex128
	mycomplex = complex(mybigfloat, mybigfloat)
	fmt.Println(mycomplex)

	var mystring string = "Anuragyadav"
	fmt.Println(mystring)
	mystring = `hello My name Is
	
	Anurag yadav`
	fmt.Println(mystring)
	var age = 10
	fmt.Println(age)
	age2 := 20
	fmt.Println(age2)
	age, name := 25, "Anurag"
	fmt.Println(age, name)

	//constants
	// typed constsnts
	// untyped constants
	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = x
	fmt.Print(b)
	b = float64(y)
	fmt.Println(b)

}

func idenfierscode() {
	var _myInt int = 10
	fmt.Println(_myInt)

	a := 10
	b := 30
	sum := a + b
	firstName := "Anurag"
	lastname := "Yadav"
	fmt.Printf("%s %s", firstName, lastname)
	fmt.Println(sum)

}

func arrayuse() {
	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

}

func slicelkaro() {
	var a []int
	fmt.Println(a == nil)
	a = []int{}
	fmt.Println(a == nil)

}
func structfun() {

	type student struct {
		firstname string
		lastName  string
		age       int
		classes   []string
	}
	foo := student{
		firstname: "Anurag",
		lastName:  "Yadav",
		age:       25,
		classes:   []string{"maths", "science", "Dsa"},
	}
	foo.firstname = "Jai"
	fmt.Println(foo)
	var s2 student

	fmt.Printf("%+v", s2)
}

func consitionalfun() {
	a := 10
	if a < 10 {
		fmt.Println(a)
	} else {
		fmt.Println("Okk")
	}

	if even := isEven(a); even {
		fmt.Println("Age is even")
	} else {
		fmt.Println("Not even")
	}
}
func isEven(a int) bool {
	a = a + 1
	return a&1 == 0
}

func shadowingjkaro() {
	x := 20
	{
		x := 30
		fmt.Println(x)
	}
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	a := "Anurag yadav"
	for idx, runeValue := range a {
		fmt.Printf("%d %c", idx, runeValue)
	}

	b := 100
	for {
		b--
		if b == 90 {
			break
		}
	}
	print("OKKKKK")

outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue outerLoop
			}
			fmt.Println(i, j)
		}
	}
	ans := saymayname()
	print(ans)

	name, age := agename()
	fmt.Println(name, age)
	sum := varidicsum(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)

}
func saymayname() string {
	return "Anurag yadav"
}

func agename() (string, int) {
	return "Anurag yadav", 20
}

func varidicsum(nums ...int) int {
	t := 0
	for _, num := range nums {
		t += num
	}
	return t
}

func double(x int) int {
	return x * 2
}

func apply(f func(int) int, x int) int {
	return f(x)
}
func closures() {
	result := apply(double, 5)
	fmt.Println(result)
	muxi := multiplier(10)
	fmt.Println(muxi(10))

}

func multiplier(mul int) func(int) int {
	return func(i int) int {
		return i * mul
	}
}
