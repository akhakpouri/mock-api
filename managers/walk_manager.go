package managers

func Walk(x interface{}, fn func(input string)) {
	fn("this really is awesome!")
}
