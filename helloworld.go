package main

import "fmt"
import (
	"./def"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	if true {
		fmt.Println("Hello, world")
	}
	var i = 0
	var t chan int
	//var j int
	b := "WTF"
	fmt.Println(&b)
	print(&i)
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(b)
	go func () {
		fmt.Println("before executed!")
		t<- 1
		fmt.Println("I'm executed")
	}()

	time.Sleep(5 * time.Second)

	select {
	case <-t:
		fmt.Println("i is less than 10")
	default:
		fmt.Println("hahahaha")

	}
	var c = &b
	fmt.Println(c)
	fmt.Println(os.Getpid())
	var intPtr *int
	intPtr = &i
	if intPtr == nil {
		fmt.Println("int_prt init value is %x", intPtr)
	} else {
		fmt.Println("Hahhahahahhh")
	}

	var slice_ = make([]int, 5, 10)
	fmt.Println(slice_)

	var map1 = make(map[string]string)
	map1["foo"] = "bar"
	map1["hello"] = "world"

	fmt.Println(len(map1), map1)

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
	fmt.Println(string(add(8, 61)))
	fmt.Println(concat("W", "W"))
	assertType(1, "hello")
	loadTest()
}

func add(x int, y int) int {
	return x + y
}

func concat(a string, b string) string {
	return a + b
}


func assertType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			//fmt.Println(index)
			fmt.Println(arg, "is int")
		default:
			fmt.Println("Nothing")
		}
	}
}

func loadTest()  {
	def.FuncTest()
}