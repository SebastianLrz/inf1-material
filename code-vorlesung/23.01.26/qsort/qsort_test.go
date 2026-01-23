package qsort

import "fmt"

func Qsort(l []int) {
	//Sonderfall: litse ist leer oder hat kein element
	if len(l) <= 1 {
		return
	}

	pivot := l[0]

	left := []int{}
	right := []int{}

	//Partitionieren der Liste
	// kleinere elemente  als das Pivot nach links, grössere nach rechts
	for _, el := range l[1:] {
		if el < pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	Qsort(left)
	Qsort(right)

	//elemente in die ursprüngliche liste zurückkopieren
	for i, el := range left {
		l[i] = el
	}

	for i, el := range right {
		l[i+len(left)+1] = el
	}

}

func ExampleQsort() {
	l1 := []int{17, 25, 22, 3, 15, 4, 35, 105, 42, 1}

	Qsort(l1)
	fmt.Println(l1)

	//Output:
	// [1 3 4 15 17 22 25 35 42 105]
}
