package test

func Channel() {
	value := make(chan int)
	defer close(value)
}
