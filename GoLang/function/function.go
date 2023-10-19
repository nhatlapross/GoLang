package main

import "fmt"

// hàm không trả giá trị
func SimpleFunction() {
	fmt.Println("Hello World")
}

//hàm trả về giá trị int
func add(x int, y int) int {
	return x + y
}

// hàm trả về giá trị string
func addString(x string, y string) string {
	return x + y
}

// hàm trả về biến được đặt tên
func rectangle(l int, b int) (parameter int) {
	//var parameter int
	parameter = 2 * (l + b)
	//fmt.Println("Parameter: ", parameter)

	//area = l * b
	return // Return statement without specify variable name
}

// hàm trả về nhiều giá trị
func rectangle_multipleValue(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

// hàm trả về nhiều giá trị
func rectangle_multipleValue2(l int, b int) (int, int) {

	return 2 * (l + b), l * b // Return statement without specify variable name
}

// hàm lưu lại thay đổi giá trị của biến truyền vào
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

// định nghĩa function là 1 biến
var (
	area = func(l int, b int) int {
		return l * b
	}
)

// higher order function
func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

// Hàm trả về hàm
func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// Có thể định nghĩa hàm theo kiểu biến
type First func(int) int
type Second func(int) First

func squareSum2(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

//Variadic Functions : hàm cho phép vô số tham số cùng kiểu
func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s)
}

func calculation(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

//variadic đa kiểu dữ liệu

func variadicMultiArg(i ...interface{}) {
	for _, v := range i {
		fmt.Printf("%v--%T \n", v, v)
	}
}

//Recursion: phép toán đệ quy
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //=>fact(3) = 3* fact(2) = 3*2*1
}

func main() {
	SimpleFunction()

	sum := add(42, 13)
	fmt.Println(sum)

	newString := addString("Alvin", "Ichi")
	fmt.Println(newString)

	fmt.Println(rectangle(3, 2))

	var a, p int
	a, p = rectangle_multipleValue(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)

	_, p = rectangle_multipleValue2(20, 30)
	//fmt.Println("Area2:", a)
	fmt.Println("Parameter2:", p)

	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)
	update(&age, &text)
	fmt.Println("After :", text, age)

	fmt.Println(area(20, 30))

	//function nặc danh
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	//Closures function
	l := 20
	b := 30
	func() {
		var area int
		area = l * b // sử dụng biến cục bộ trong hàm
		fmt.Println(area)
	}()

	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	//Hàm partialSum sẽ trả về hàm sum (3,y)
	partial := partialSum(3)
	fmt.Println(partial(7))

	//Hàm trả về hàm
	fmt.Println(squareSum(5)(6)(7))
	fmt.Println(squareSum2(5)(6)(7))

	//variadic
	variadicExample("red", "blue", "green", "yellow", "black")

	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 20))

	//variadic đa tham số đa kiểu dữ liệu
	variadicMultiArg(1, "red", true, 10.5, []string{"foo", "bar", "baz"}, map[string]int{"apple": 23, "tomato": 13})

	// đệ quy sẽ gọi lại chính hàm đó cho đến khi thỏa dkien để thoát khỏi vòng lặp
	fmt.Println(fact(7))

	// ví dụ dãy fibonaci sử dụng đệ quy
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(3))

}
