package sort

import (
	"fmt"
	"sort"
)

type DataManager interface {
	Sorting()
}

type Str struct {
	V   []string
	ASC bool
}

type Itr struct {
	V   []int
	ASC bool
}

func (t *Str) Sorting() {
	sort.Slice(t.V, func(i, j int) bool {
		if t.ASC {
			return t.V[i] < t.V[j]
		}
		return t.V[i] > t.V[j]
	})
	fmt.Println(t.V)
}

func (t *Itr) Sorting() {
	sort.Slice(t.V, func(i, j int) bool {
		if t.ASC {
			return t.V[i] < t.V[j]
		}
		return t.V[i] > t.V[j]
	})
	fmt.Println(t.V)
}

// func main() {
// 	var svc DataManager

// 	svc = &Str{
// 		v:   []string{"a", "f", "d", "r", "t", "e"},
// 		asc: true,
// 	}
// 	svc.Sorting()

// 	svc = &Itr{v: []int{6, 4, 7, 3, 5, 2}}
// 	svc.Sorting()
// }
