package main

func myfuncx1() []MyStruct {
	return []MyStruct{MyStruct{Val: 1}}
}

func myfuncx2() []*MyStruct {
	return []MyStruct{&MyStruct{Val: 1}}
}

func myfuncx3(s *[]MyStruct) {
	*s = []MyStruct{MyStruct{Val: 1}}
}

func myfuncx4(s *[]*MyStruct) {
	*s = []MyStruct{&MyStruct{Val: 1}}
}

func main() {

}
