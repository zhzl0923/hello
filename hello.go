package hello

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func SayHello() {
	str := Hello()
	fmt.Println(str)
}
