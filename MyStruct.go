package main

type MyStruct struct {
	Val int
}

func myfunc1() MyStruct {
	return MyStruct{Val: 1}
}

func myfunc2() *MyStruct {
	return &MyStruct{}
}

func myfunc3(s *MyStruct) {
	s.Val = 1
}

func main() {

}
