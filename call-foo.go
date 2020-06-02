package main

func foo([]interface{}) {

}

func main() {
	var a []string = []string{"hello", "world"}
	// can not call directly
	//foo(a)
	b := make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	foo(b)
}
