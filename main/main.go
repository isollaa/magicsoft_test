package main

import "fmt"

import "github.com/isollaa/magicsoft_test/lala"

type Comparator interface {
	Compare(int, int) bool
}

type asc struct{}

func (t *asc) Compare(before, current int) bool {
	return before > current
}

type desc struct{}

func (t *desc) Compare(before, current int) bool {
	return current > before
}

func BubbleSort(numbers []int, c Comparator) []int {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if c.Compare(numbers[j-1], numbers[j]) {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	return numbers
}

func main() {
	var svc lala.DataManager

	svc = &lala.Str{
		V:   []string{"a", "f", "d", "r", "t", "e"},
		ASC: false,
	}
	svc.Sorting()
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(BubbleSort(a, &asc{}))
	fmt.Println(BubbleSort(a, &desc{}))
}
