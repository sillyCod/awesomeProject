package main

import "fmt"
import "os"

func main()  {
	fmt.Println("Hello, world!")
	if true {
		fmt.Println("Hello, world")
	}
	var i = 0
	var t chan int
	var j int
	b := "WTF"
	fmt.Println(&b)
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(b)

	select {
	case t <- j:
		fmt.Println("i is less than 10")
	default:
		fmt.Println("hahahaha")

	}
	var c *string = &b
	fmt.Println(c)
	fmt.Println(os.Getpid())
	var intPtr *int
	if intPtr == nil {
		fmt.Println("int_prt init value is %x", intPtr)
	}

	var slice_ []int = make([]int, 5, 10)
	fmt.Println(slice_)

	var map1 map[string]string = make(map[string]string)
	map1["foo"] = "bar"
	map1["hello"] = "world"

	fmt.Println(map1)
	const (
		num1 = iota
		num2
		num3
		num4 = "Hello"
		num5
		num6 = iota

	)
	fmt.Println(num1, num2, num3, num4, num5, num6)
}