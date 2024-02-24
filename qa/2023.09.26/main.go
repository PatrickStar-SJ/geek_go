package main

import "fmt"

func main() {
	//DeferClosureLoopV1()
	//DeferClosureLoopV2()
	//DeferClosureLoopV3()
	OfNullable[User](User{}).Apply(func(t User) {
		println(t.Name)
	})
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("循环 %p \n", &i)
		defer func() {
			fmt.Printf("%p \n", &i)
			println(i)
		}()
	}
	println("跳出循环")
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("%p \n", &val)
			println(val)
		}(i)
	}
	println("跳出循环")
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("%p \n", &j)
			println(j)
		}()
	}
	println("跳出循环")
}
