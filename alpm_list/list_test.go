package alpm_list

import "fmt"
import "unsafe"

func ExampleList() {
	var l *List
	var strl *List

	ints := []int{1, 2, 3, 4, 5}

	for k := range ints {
		Append(&l, uintptr(unsafe.Pointer(&ints[k])))
	}
	for i := l; i != nil; i = i.Next() {
		data := (*int)(unsafe.Pointer(i.Data()))
		fmt.Println(*data)
	}
	fmt.Println("Size:", l.Count())

	AppendStrdup(&strl, "a")
	AppendStrdup(&strl, "b")
	AppendStrdup(&strl, "c")
	AppendStrdup(&strl, "d")
	AppendStrdup(&strl, "e")
	AppendStrdup(&strl, "f")
	for i := strl; i != nil; i = i.Next() {
		data := i.String()
		fmt.Println(data)
	}
	strl = strl.Reverse()
	for i := strl; i != nil; i = i.Next() {
		data := i.String()
		fmt.Println(data)
	}
	fmt.Println("Size:", strl.Count())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// Size: 5
	// a
	// b
	// c
	// d
	// e
	// f
	// f
	// e
	// d
	// c
	// b
	// a
	// Size: 6
}
