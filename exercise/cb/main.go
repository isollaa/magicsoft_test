package main

import "fmt"
import "unsafe"
import "strconv"

func castStr(v *string) string {
	return fmt.Sprint(uintptr(unsafe.Pointer(v)))
}

func uncastStr(s string) string {
	p, _ := strconv.ParseInt(s, 10, 64)
	return *((*string)(unsafe.Pointer(uintptr(p))))
}

func main() {

	// a := []int{5, 4, 6, 3, 7, 2, 8, 1, 0}
	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }
	// fmt.Print(a)

	b := 4
	switch {
	case 3 > b, b == 5: // can use || && ,
		println("ok", b)
	default:
		println("emm", b)
	}

	// onevar := "something"
	// other := "something else"
	// sa := []string{castStr(&onevar), castStr(&other)}

	// for _, v := range sa {
	//     fmt.Printf("{{%s}}\n", v)
	//     fmt.Printf("%v\n", uncastStr(v))
	// }

	//for _, v := range sa {
	//  vName := fmt.Sprintf("{{%s}}", v)
	//  msg = strings.Replace(msg, vName, uncastStr(v) -1)
	//}
}
