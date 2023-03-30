package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"
	fmt.Println("new:", s)
	fmt.Println("newindex:", s[1])

	fmt.Println(len(s))

	s = append(s, "four")
	s = append(s, "five", "six")
	fmt.Println(s)

}
