package main

var cache map[int]uint64 = make(map[int]uint64)

func fibonacci(n int) uint64 {

	var val uint64

	if val, ok := cache[n]; ok {
		return val
	}

	if n == 1 {
		val = 1
	} else if n == 2 {
		val = 1
	} else if n > 2 {
		val = fibonacci(n-1) + fibonacci(n-2)
	}

	cache[n] = val
	return val
}

func main() {
	for i := 0; i < 100000; i++ {
		println(fibonacci(i))
	}
}
