package test

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func Test_1(t *testing.T) {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	fmt.Println(newSlice, len(newSlice), cap(newSlice))
	newSlice = append(newSlice, 100)
	fmt.Println(slice)

	tSlice := [...]int{1,2,4:5,6}
	//tSlice := [...]int{1,2,0:5,6}
	fmt.Println(tSlice)
}

func Test_2(t *testing.T) {
	fmt.Println(strings.TrimRight("abcdcba", "abc"))
	fmt.Println(strings.TrimRight("abcdaaa", "abc"))
	fmt.Println(strings.TrimRight("baba", "ba"))
	fmt.Println(strings.TrimRight("baba", "ab"))
	fmt.Println(strings.TrimRight("baba", "ba"))
	fmt.Println(strings.TrimRight("baba", "a"))
}

type test1 struct {

}
type test2 struct {
	a int64
}
type test3 struct {
	a int8
	s string
}
type test4 struct {
	a int64
	s string
	s1 *string
}
func Test_3(t *testing.T) {
	fmt.Println(unsafe.Sizeof(test1{}))
	fmt.Println(unsafe.Sizeof(test2{}))
	fmt.Println(unsafe.Sizeof(test3{}))
	fmt.Println(unsafe.Sizeof(test4{}))
}

func Test_5(t *testing.T) {
	test1 := []int{1,2,3}
	ttt1(test1...)
	//ttt1(test1)

	test2 := []interface{}{1,2,3}
	ttt2(test2)
	ttt2(test2...)

}

func ttt1(a ...int) {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
func ttt2(a ...interface{}) {
	fmt.Println(a)
	fmt.Println(a...)
	fmt.Printf("%T\n", a)
}

func Test_6(t *testing.T) {
	//reflect
}