package examples

type MyObject struct{}

func ExampleOne() {
	myThing := &MyObject{}
	// do something with myThing
}

func ExampleTwo(thing *MyObject) {
	// do something with myThing
}
