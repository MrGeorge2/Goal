package goalist

// Iterate over the goalist and call the function from parameter on each one
func (l Goalist[T]) ForEach(function func(x T)) {
	for _, item := range l {
		function(item)
	}
}

// Iterate over the goalist and call the function from parameter on each one
func (l Goalist[T]) GoForEach(function func(x T)) {
	for _, item := range l {
		go function(item)
	}
}
