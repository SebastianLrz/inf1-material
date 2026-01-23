package slicing

import "fmt"

func Example() {
	l1 := []int{11, 22, 33, 44, 55}
	s1 := l1[2:4]

	fmt.Println(l1)
	fmt.Println(s1)

	s1[0] = 555

	fmt.Println(l1)
	fmt.Println(s1)

	s1 = append(s1, 777)

	fmt.Println(l1)
	fmt.Println(s1)

	// Ouput:
	// []
}
